package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/builtin"
)

func execCmd(shellCmd string, shellArgs []string) {
	if cmdFunc, ok := builtin.CmdFuncMap[shellCmd]; ok == true {
		cmdFunc.Execute(shellArgs)
	} else {
		if path, _ := exec.LookPath(shellCmd); path != "" {
			cmd := exec.Command(shellCmd, shellArgs...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			fmt.Println(shellCmd + ": command not found")
		}
	}
}
