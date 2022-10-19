package main

import (
	"os"
	"os/exec"
)

const (
	ExitCodeIOError         int = 17
	ExitCodeCommandNotFound int = 111
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) < 1 {
		return ExitCodeCommandNotFound
	}
	command, args := cmd[0], cmd[1:]
	cmdResult := exec.Command(command, args...) //nolint
	cmdResult.Env = env.Strings()
	cmdResult.Stdin = os.Stdin
	cmdResult.Stdout = os.Stdout
	cmdResult.Stderr = os.Stderr

	if err := cmdResult.Run(); err != nil {

		return ExitCodeIOError
	}
	return 0
}
