package builder

import (
	"errors"

	"github.com/evanw/esbuild/pkg/api"
)

type (
	CompiledResult struct {
		Source []byte
		Path   string
		Hash   string
	}

	BuildResult struct {
		Compiled CompiledResult
		Errors   []error
	}

	BuildOptions struct {
		WorkingFolder string
		EntryPoint    string
	}

	WorkflowBuilder interface {
		Build() BuildResult
	}

	builder struct {
		buildOptions BuildOptions
	}
)

func NewWorkflowBuilder(options BuildOptions) WorkflowBuilder {
	return &builder{
		buildOptions: options,
	}
}

func (b *builder) Build() BuildResult {
	opts := api.BuildOptions{
		Loader: map[string]api.Loader{
			".js":  api.LoaderJS,
			".tsx": api.LoaderTSX,
			".ts":  api.LoaderTS,
			".css": api.LoaderCSS,
		},
		AbsWorkingDir:    b.buildOptions.WorkingFolder,
		Target:           api.ESNext,
		Format:           api.FormatCommonJS,
		Sourcemap:        api.SourceMapInline,
		SourcesContent:   api.SourcesContentInclude,
		LegalComments:    api.LegalCommentsNone,
		Platform:         api.PlatformDefault,
		LogLevel:         api.LogLevelError,
		Charset:          api.CharsetUTF8,
		EntryPoints:      []string{b.buildOptions.EntryPoint},
		Outdir:           "dist",
		Bundle:           true,
		Write:            false,
		TreeShaking:      api.TreeShakingTrue,
		MinifyWhitespace: true,
		MinifySyntax:     true,
	}
	tr := api.Build(opts)

	result := BuildResult{}

	if len(tr.OutputFiles) > 0 {

		if len(tr.Errors) > 1 {
			result.addError(errors.New("build produced multiple files, a single output is supported only"))
		}

		file := tr.OutputFiles[0]
		result.Compiled = CompiledResult{
			Source: file.Contents,
			Path:   file.Path,
			Hash:   file.Hash,
		}
	}

	return result
}

func (br *BuildResult) HasOutput() bool {
	return len(br.Compiled.Source) > 0
}

func (br *BuildResult) addError(err error) {
	br.Errors = append(br.Errors, err)
}
