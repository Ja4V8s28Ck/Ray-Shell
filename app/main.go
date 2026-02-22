package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/builtin"
)

var historyArrPtr int

func main() {
	for {
		cmdLine, err := readLine()

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if cmdLine != "" {
			builtin.StoreHistory(cmdLine)
			historyArrPtr = builtin.HistoryArrCount
		}

		shellCmd, shellArgs := parseCmd(cmdLine)
		if shellCmd == "" {
			continue
		}

		execCmd(shellCmd, shellArgs)
	}
}
