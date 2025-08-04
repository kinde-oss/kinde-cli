package cmd

import (
	"fmt"
	"reflect"
	"time"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-go/kinde"
	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type manageCmd struct {
	cmd *cobra.Command
}

func newManageCmd() *manageCmd {

	manageCmd := &manageCmd{}

	manageCmd.cmd = &cobra.Command{
		Use:   "manage",
		Args:  cobra.NoArgs,
		Short: "Manage Kinde business",
		RunE:  manageCmd.runLogin,
	}

	business := &cobra.Command{
		Use:  "business",
		Args: cobra.MatchAll(),
	}

	manageCmd.cmd.AddCommand(business)

	addCommand[management_api.UpdateBusinessReq,
		management_api.SuccessResponse,
		management_api.UpdateBusinessTooManyRequests,
		management_api.UpdateBusinessForbidden,
		management_api.UpdateBusinessBadRequest,
	](business, &cobra.Command{
		Use:   "update",
		Args:  cobra.ArbitraryArgs,
		Short: "Update business information",
		Long:  "Update information about the Kinde business, such as name, domain, and other details.",
	}, "UpdateBusiness")
	addCommand[management_api.UpdateBusinessReq,
		management_api.GetBusinessResponse,
		management_api.GetBusinessTooManyRequests,
		management_api.GetBusinessForbidden,
		management_api.GetBusinessBadRequest,
	](business, &cobra.Command{
		Use:   "get",
		Args:  cobra.ArbitraryArgs,
		Short: "Get business information",
		Long:  "Retrieve information about the Kinde business, such as name, domain, and other details.",
	}, "GetBusiness")

	return manageCmd
}

func addCommand[T, TResponse, TThrottledResponse, TForbiddenResponse, TBadRequestResponse any](parentCommand, command *cobra.Command, oasClientMethodName string) {
	generateCommandFlags[T](command)

	managementApiResourceCommand := createManagementApiCommand[TResponse, TThrottledResponse, TForbiddenResponse, TBadRequestResponse](
		command,
		func(apiClient *management_api.Client, cmd *cobra.Command) (response any, err error) {
			instance := mapFlagsToStruct[T](cmd.Flags())
			method := reflect.ValueOf(apiClient).MethodByName(oasClientMethodName)
			numArgs := method.Type().NumIn()
			args := []reflect.Value{reflect.ValueOf(cmd.Context()), reflect.ValueOf(instance)}
			result := method.Call(args[0:numArgs])
			res := result[0].Interface()
			errInterface := result[1].Interface()
			log.Info().Any("response", res).Msg("API call response")
			if errInterface == nil {
				return res, nil
			}
			return res, errInterface.(error)
		})
	parentCommand.AddCommand(managementApiResourceCommand)
}

func generateCommandFlags[T any](command *cobra.Command) {
	instance := new(T)
	t := reflect.TypeOf(*instance)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")

		switch field.Type {
		case reflect.TypeOf(management_api.OptNilBool{}), reflect.TypeOf(management_api.OptBool{}):
			command.Flags().Bool(tag, false, "")
		default:
			command.Flags().String(tag, "", "")
		}
	}
}

func mapFlagsToStruct[T any](flagSet *pflag.FlagSet) *T {
	instance := new(T)
	t := reflect.TypeOf(*instance)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")

		switch field.Type {
		case reflect.TypeOf(management_api.OptNilString{}):
			flag := flagSet.Lookup(tag)
			if flag.Changed {
				val := reflect.ValueOf(instance).Elem()
				val.Field(i).Set(reflect.ValueOf(management_api.NewOptNilString(flag.Value.String())))
			}
		case reflect.TypeOf(management_api.OptNilBool{}):
			flag := flagSet.Lookup(tag)
			if flag.Changed {
				val := reflect.ValueOf(instance).Elem()
				if flagValue, err := flagSet.GetBool(tag); err == nil {
					val.Field(i).Set(reflect.ValueOf(management_api.NewOptNilBool(flagValue)))
				}
			}
		default:
			if value, err := flagSet.GetString(tag); err == nil {
				val := reflect.ValueOf(instance).Elem()
				val.Field(i).Set(reflect.ValueOf(management_api.NewOptNilString(value)))
			}
		}
	}
	return instance
}

func createManagementApiCommand[TResponse, TThrottledResponse, TForbiddenResponse, TBadRequestResponse any](cmd *cobra.Command, f func(apiClient *management_api.Client, cmd *cobra.Command) (response any, err error)) *cobra.Command {
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		log := log.Ctx(cmd.Context())
		config := config.FromContext[config.Config](cmd.Context())

		env := config.GetEnvironment()

		if env.DomainName == "" {
			return fmt.Errorf("no environment configured. Please run 'kinde login'")
		}

		clientCredentials, err := env.NewClientCredentialsFlow()
		if err != nil {
			return fmt.Errorf("failed to create client credentials flow: %w", err)
		}

		kindeDomainUrl := fmt.Sprintf("https://%s", env.DomainName)

		managementApi, err := kinde.NewManagementAPI(cmd.Context(), kindeDomainUrl, clientCredentials)
		if err != nil {
			return fmt.Errorf("failed to create management API client: %w", err)
		}

		res, errResponse, badRequest, err := apiCallAs[TResponse, TThrottledResponse, TForbiddenResponse, TBadRequestResponse](f, managementApi, cmd)
		if err != nil {
			return fmt.Errorf("failed to get business: %w", err)
		}
		if errResponse != nil {
			log.Error().Any("error", errResponse).Msg("Failed calling management API")
			return fmt.Errorf("failed to call API")
		}
		if badRequest != nil {
			log.Error().Any("bad_request", badRequest).Msg("Bad request to management API")
			return fmt.Errorf("bad request: %s", badRequest)
		}
		log.Debug().Any("api_response", res).Msg("API call")
		return nil
	}

	return cmd
}

func (c *manageCmd) runLogin(cmd *cobra.Command, args []string) error {

	return nil
}

func apiCallAs[TResponse, TThrottledResponse, T1, T2 any](f func(apiClient *management_api.Client, cmd *cobra.Command) (response any, err error), apiClient *management_api.Client, cmd *cobra.Command) (*TResponse, *T1, *T2, error) {
	var response any
	var err error

	i := 0
	for {
		response, err = f(apiClient, cmd)
		i++

		if i > 10 {
			return nil, nil, nil, fmt.Errorf("too many retries: %w", err)
		}

		if _, ok := response.(*TThrottledResponse); ok {
			time.Sleep(1 * time.Second) // Wait before retrying
			continue
		}
		break
	}

	if err != nil {
		return nil, nil, nil, err
	}
	if res, ok := response.(*TResponse); ok {
		return res, nil, nil, nil
	}
	if res, ok := response.(TResponse); ok {
		return &res, nil, nil, nil
	}
	if res, ok := response.(*T1); ok {
		return nil, res, nil, nil
	}
	if res, ok := response.(T1); ok {
		return nil, &res, nil, nil
	}
	if res, ok := response.(*T2); ok {
		return nil, nil, res, nil
	}
	if res, ok := response.(T2); ok {
		return nil, nil, &res, nil
	}

	return nil, nil, nil, fmt.Errorf("Failed to cast response, received type: %s", reflect.TypeOf(response).String())
}
