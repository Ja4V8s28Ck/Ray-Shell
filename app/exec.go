package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/builtin"
)

func execCmd(shellCmd string, shellArgs []string) {
	n := len(shellArgs)

	var stdout *os.File = os.Stdout
	if n > 2 && (shellArgs[n-2] == ">" || shellArgs[n-2] == "1>") {
		outputFile, err := os.Create(shellArgs[n-1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer outputFile.Close()
		stdout = outputFile
		shellArgs = shellArgs[:n-2]
	}

	if cmdFunc, ok := builtin.CmdFuncMap[shellCmd]; ok {
		ctx := &builtin.ExecContext{
			Stdin:  os.Stdin,
			Stdout: stdout,
			Stderr: os.Stderr,
		}

		cmdFunc.Execute(shellArgs, ctx)
	} else {
		if path, _ := exec.LookPath(shellCmd); path != "" {
			cmd := exec.Command(shellCmd, shellArgs...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			fmt.Println(shellCmd + ": command not found")
		}
	}
}
