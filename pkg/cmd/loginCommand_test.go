package cmd

import (
	"testing"

	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/stretchr/testify/assert"
)

func TestNewLoginCmd_ReturnsLoginCmd(t *testing.T) {
	ios, _, _, _ := iostreams.Test()
	cmdStruct := newLoginCmd(ios)
	assert.NotNil(t, cmdStruct)
	assert.NotNil(t, cmdStruct.cmd)
}

func TestNewLoginCmd_CommandProperties(t *testing.T) {
	ios, _, _, _ := iostreams.Test()
	cmdStruct := newLoginCmd(ios)
	cmd := cmdStruct.cmd

	assert.Equal(t, "login", cmd.Use)
	assert.Equal(t, "Log in to use the Management API, through the browser or with an M2M application", cmd.Short)
	assert.NotNil(t, cmd.RunE)
}

func TestNewLoginCmd_RunEIsSet(t *testing.T) {
	ios, _, _, _ := iostreams.Test()
	cmdStruct := newLoginCmd(ios)
	assert.NotNil(t, cmdStruct.cmd.RunE)
}
