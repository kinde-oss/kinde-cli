package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewLogoutCmd_ReturnsLogoutCmd(t *testing.T) {
	cmd := newLogoutCmd()
	assert.NotNil(t, cmd, "newLogoutCmd should not return nil")
	assert.NotNil(t, cmd.cmd, "logoutCmd.cmd should not be nil")
	assert.Equal(t, "logout", cmd.cmd.Use, "Command 'Use' should be 'logout'")
	assert.Equal(t, "Logout and clear authentication tokens", cmd.cmd.Short, "Command 'Short' should match")
	assert.IsType(t, &cobra.Command{}, cmd.cmd, "cmd.cmd should be of type *cobra.Command")
	assert.NotNil(t, cmd.cmd.RunE, "RunE should be set")
}
