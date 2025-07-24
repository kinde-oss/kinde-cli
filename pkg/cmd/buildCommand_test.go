package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newBuildCmd_CommandProperties(t *testing.T) {
	bc := newBuildCmd()
	cmd := bc.cmd

	assert.Equal(t, "build", cmd.Use)
	assert.Equal(t, "Creating deployable version of code in ./kinde/output", cmd.Short)
	assert.NotNil(t, cmd.Args)
	assert.NotNil(t, cmd.Run)
}

func Test_newBuildCmd_Flags(t *testing.T) {
	bc := newBuildCmd()
	cmd := bc.cmd

	flag := cmd.Flags().Lookup("cwd")
	assert.NotNil(t, flag)
	assert.Equal(t, "The directory to build from", flag.Usage)
	assert.Equal(t, "", flag.Value.String())
}

func Test_newBuildCmd_CwdFlagAssignment(t *testing.T) {
	bc := newBuildCmd()
	cmd := bc.cmd

	_ = cmd.Flags().Set("cwd", "/tmp/testdir")
	assert.Equal(t, "/tmp/testdir", bc.cwd)
}

func Test_newBuildCmd_RunFunctionExists(t *testing.T) {
	bc := newBuildCmd()
	assert.NotNil(t, bc.cmd.Run)
}
