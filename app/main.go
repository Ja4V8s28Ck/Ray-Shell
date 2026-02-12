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

func main() {
	// Map commands to it's function
	cmdFuncMap := map[string]func(shellArgs []string){
		"exit": func(shellArgs []string) { os.Exit(0) },
		"echo": func(shellArgs []string) { fmt.Println(strings.Join(shellArgs, " ")) },
	}

	buffReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		cmdLine, err := buffReader.ReadString('\n')
		if err != nil {
			logFatal(err)
		}

		var shellArgs []string
		cmdLineArr := strings.Split(cmdLine[:len(cmdLine)-1], " ")
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
