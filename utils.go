package main

import (
	"os"
	"os/exec"
	"runtime"
)
// clearTerminal use cmd commands to clean the screen
func clearTerminal() {

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
