package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		cmdLine, err := readLine("$")

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		shellCmd, shellArgs := parseCmd(cmdLine)
		if shellCmd == "" {
			continue
		}

		execCmd(shellCmd, shellArgs)
	}
}
