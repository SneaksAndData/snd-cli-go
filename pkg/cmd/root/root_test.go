package root

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmdRoot(t *testing.T) {
	cmd, err := NewCmdRoot()
	assert.NoError(t, err, "NewCmdRoot should not return an error")

	// Capture the output
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	// Simulate executing the command without any subcommands and flags
	err = cmd.Execute()
	assert.Error(t, err, "unknown command \"^\\\\QTestNewCmdRoot\\\\E$\" for \"snd\"")

	// Simulate executing the command with -h flag
	cmd.SetArgs([]string{"-h"})
	err = cmd.Execute()
	assert.NoError(t, err, "Executing root command should not return an error")

}
