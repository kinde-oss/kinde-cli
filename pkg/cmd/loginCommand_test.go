package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLoginCmd_ReturnsLoginCmd(t *testing.T) {
	cmdStruct := newLoginCmd()
	assert.NotNil(t, cmdStruct)
	assert.NotNil(t, cmdStruct.cmd)
}

func TestNewLoginCmd_CommandProperties(t *testing.T) {
	cmdStruct := newLoginCmd()
	cmd := cmdStruct.cmd

	assert.Equal(t, "login", cmd.Use)
	assert.Equal(t, "Login to Kinde", cmd.Short)
	assert.Nil(t, cmd.Args)
	assert.NotNil(t, cmd.RunE)
}

func TestNewLoginCmd_RunEIsSet(t *testing.T) {
	cmdStruct := newLoginCmd()
	assert.NotNil(t, cmdStruct.cmd.RunE)
}
