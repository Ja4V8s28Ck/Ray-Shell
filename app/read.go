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
	cmdLine = strings.TrimRight(cmdLine, "\n")
	cmdLineArr := parseArgs(cmdLine)
	shellCmd := cmdLineArr[0]

	if len(cmdLineArr) > 1 {
		shellArgs = cmdLineArr[1:]
	}

	return shellCmd, shellArgs
}

func parseArgs(shellArgString string) []string {
	var args []string
	var stringBuilder strings.Builder
	var inQuote rune // 0 means not in quote, '"' or '\'' means in that quote
	n := len(shellArgString)

	for i := 0; i < n; i++ {
		char := rune(shellArgString[i])

		switch {
		case char == '"' || char == '\'':
			if inQuote == 0 {
				inQuote = char
			} else if inQuote == char {
				inQuote = 0
			} else {
				stringBuilder.WriteRune(char)
			}

		case char == ' ' && inQuote == 0:
			if stringBuilder.Len() > 0 {
				args = append(args, stringBuilder.String())
				stringBuilder.Reset()
			}

		case char == '\\' && i+1 < n:
			if inQuote == '\'' {
				stringBuilder.WriteRune(char)
			}
			stringBuilder.WriteRune(rune(shellArgString[i+1]))
			i++

		default:
			stringBuilder.WriteRune(char)
		}
	}

	if stringBuilder.Len() > 0 {
		args = append(args, stringBuilder.String())
	}

	return args
}
