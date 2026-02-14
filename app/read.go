package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readCmd(buffReader *bufio.Reader) (string, []string) {
	cmdLine, err := buffReader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	var shellArgs []string
	cmdLineArr := strings.Fields(cmdLine[:len(cmdLine)-1])
	shellCmd := cmdLineArr[0]
	if len(cmdLineArr) > 1 {
		shellArgs = cmdLineArr[1:]
	}

	return shellCmd, shellArgs
}
