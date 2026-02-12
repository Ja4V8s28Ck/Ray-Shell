package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func logFatal(err error) {
	fmt.Fprintln(os.Stderr, "Error reading input:", err)
	os.Exit(1)
}

func invalidCommand(shellCmd string) {
	fmt.Println(shellCmd + ": command not found")
}

func main() {
	// Map commands to it's function
	var cmdFuncMap map[string]func(shellArgs []string)
	cmdFuncMap = map[string]func(shellArgs []string){
		"echo": func(shellArgs []string) { fmt.Println(strings.Join(shellArgs, " ")) },
		"exit": func(shellArgs []string) { os.Exit(0) },
		"type": func(shellArgs []string) {
			if len(shellArgs) == 0 {
				fmt.Println("type command needs argument")
			} else if len(shellArgs) == 1 {
				if _, ok := cmdFuncMap[shellArgs[0]]; ok {
					fmt.Println(shellArgs[0] + " is a shell builtin")
				} else {
					fmt.Println(shellArgs[0] + ": not found")
				}
			}
		},
	}

	buffReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		cmdLine, err := buffReader.ReadString('\n')
		if err != nil {
			logFatal(err)
		}

		var shellArgs []string
		cmdLineArr := strings.Fields(cmdLine[:len(cmdLine)-1])
		shellCmd := cmdLineArr[0]
		if len(cmdLineArr) > 1 {
			shellArgs = cmdLineArr[1:]
		}

		if cmdFunc, ok := cmdFuncMap[shellCmd]; ok == true {
			cmdFunc(shellArgs)
		} else {
			fmt.Println(shellCmd + ": command not found")
		}
	}
}
