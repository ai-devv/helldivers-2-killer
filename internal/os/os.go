package os

import (
	"fmt"
	"os/exec"
	"runtime"
)

func KillProcessByName(name string) error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("taskkill", "/im", name, "/F", "/T")
	} else {
		cmd = exec.Command("pkill", name)
	}

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("failed to kill process %s: %v, output: %s", name, err, output)
	}

	return nil
}
