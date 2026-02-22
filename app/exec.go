package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/builtin"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

func splitPipeline(shellCmd string, shellArgs []string) [][]string {
	var commands [][]string
	current := []string{shellCmd}
	for _, arg := range shellArgs {
		if arg == "|" {
			if len(current) > 0 {
				commands = append(commands, current)
				current = nil
			}
		} else {
			current = append(current, arg)
		}
	}
	if len(current) > 0 {
		commands = append(commands, current)
	}
	return commands
}

func execCmd(shellCmd string, shellArgs []string) {
	pipeline := splitPipeline(shellCmd, shellArgs)
	if len(pipeline) > 1 {
		execPipeline(pipeline)
		return
	}

	n := len(shellArgs)

	var stdout io.Writer = os.Stdout
	var stderr io.Writer = os.Stderr

	if n >= 2 && (utils.IsRedirectOutput(shellArgs[n-2])) {
		fileName := shellArgs[n-1]
		var outputFile *os.File

		switch shellArgs[n-2] {
		case ">", "1>":
			outputFile = utils.CreateFile(fileName)
			defer outputFile.Close()
			stdout = outputFile
		case "2>":
			outputFile = utils.CreateFile(fileName)
			defer outputFile.Close()
			stderr = outputFile
		case ">>", "1>>":
			outputFile = utils.ReadFile(fileName)
			defer outputFile.Close()
			stdout = outputFile
		case "2>>":
			outputFile = utils.ReadFile(fileName)
			defer outputFile.Close()
			stderr = outputFile
		}

		shellArgs = shellArgs[:n-2]
	}

	execSingle(shellCmd, shellArgs, os.Stdin, stdout, stderr)
}

func execSingle(shellCmd string, shellArgs []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) {
	if cmdFunc, ok := builtin.CmdFuncMap[shellCmd]; ok {
		ctx := &builtin.ExecContext{
			Stdin:  stdin,
			Stdout: stdout,
			Stderr: stderr,
		}
		cmdFunc.Execute(shellArgs, ctx)
	} else {
		if path, _ := exec.LookPath(shellCmd); path != "" {
			cmd := exec.Command(shellCmd, shellArgs...)
			cmd.Stdin = stdin
			cmd.Stdout = stdout
			cmd.Stderr = stderr
			cmd.Run()
		} else {
			fmt.Fprintln(stderr, shellCmd+": command not found")
		}
	}
}

func execPipeline(commands [][]string) {
	var input io.Reader = os.Stdin

	for i, cmdArgs := range commands {
		if i == len(commands)-1 {
			execSingle(cmdArgs[0], cmdArgs[1:], input, os.Stdout, os.Stderr)
		} else {
			pr, pw := io.Pipe()
			go func(cmdArgs []string, input io.Reader, pw *io.PipeWriter) {
				execSingle(cmdArgs[0], cmdArgs[1:], input, pw, os.Stderr)
				pw.Close()
			}(cmdArgs, input, pw)
			input = pr
		}
	}
}
