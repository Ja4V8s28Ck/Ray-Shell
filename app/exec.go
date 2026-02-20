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

	if n >= 2 && (isRedirectOutput(shellArgs[n-2])) {

		fileName := shellArgs[n-1]
		var outputFile *os.File

		switch shellArgs[n-2] {
		case ">", "1>":
			outputFile = createFile(fileName)
			defer outputFile.Close()
			stdout = outputFile
		case "2>":
			outputFile = createFile(fileName)
			defer outputFile.Close()
			stderr = outputFile
		case ">>", "1>>":
			outputFile = readFile(fileName)
			defer outputFile.Close()
			stdout = outputFile
		case "2>>":
			outputFile = readFile(fileName)
			defer outputFile.Close()
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
