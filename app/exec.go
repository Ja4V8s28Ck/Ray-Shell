package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/builtin"
)

func execCmd(shellCmd string, shellArgs []string) {
	n := len(shellArgs)

	var stdout io.Writer = os.Stdout
	var stderr io.Writer = os.Stderr

	if n > 2 && (shellArgs[n-2] == ">" || shellArgs[n-2] == "1>" || shellArgs[n-2] == "2>") {
		outputFile, err := os.Create(shellArgs[n-1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer outputFile.Close()

		switch shellArgs[n-2] {
		case ">", "1>":
			stdout = outputFile
		case "2>":
			stderr = outputFile
		}

		shellArgs = shellArgs[:n-2]
	}

	if cmdFunc, ok := builtin.CmdFuncMap[shellCmd]; ok {
		ctx := &builtin.ExecContext{
			Stdin:  os.Stdin,
			Stdout: stdout,
			Stderr: stderr,
		}

		cmdFunc.Execute(shellArgs, ctx)
	} else {
		if path, _ := exec.LookPath(shellCmd); path != "" {
			cmd := exec.Command(shellCmd, shellArgs...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = stdout
			cmd.Stderr = stderr
			cmd.Run()
		} else {
			fmt.Println(shellCmd + ": command not found")
		}
	}
}
