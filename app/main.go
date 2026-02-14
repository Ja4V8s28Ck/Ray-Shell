package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func logFatal(err error) {
	fmt.Fprintln(os.Stderr, "Error reading input:", err)
	os.Exit(1)
}

func main() {
	// Map commands to it's function
	var cmdFuncMap map[string]func(shellArgs []string)
	cmdFuncMap = map[string]func(shellArgs []string){
		"cd": func(shellArgs []string) {
			if len(shellArgs) == 1 { // Only 1 arg allowed

				if shellArgs[0] == "~" || strings.HasPrefix(shellArgs[0], "~/") { // convert relative path to abs path
					shellArgs[0] = strings.Replace(shellArgs[0], "~", os.Getenv("HOME"), 1)
				}
				if err := os.Chdir(shellArgs[0]); err != nil {
					fmt.Printf("cd: %s: No such file or directory\n", shellArgs[0])
				}

			} else {
				fmt.Println("cd: too many arguments")
			}
		},
		"echo": func(shellArgs []string) { fmt.Println(strings.Join(shellArgs, " ")) },
		"exit": func(shellArgs []string) { os.Exit(0) },
		"pwd":  func(shellArgs []string) { pwd, _ := os.Getwd(); fmt.Println(pwd) },
		"type": func(shellArgs []string) {
			if len(shellArgs) == 0 {
				fmt.Println("type: needs an argument")

			} else if len(shellArgs) == 1 {

				if _, ok := cmdFuncMap[shellArgs[0]]; ok {
					fmt.Println(shellArgs[0] + " is a shell builtin")

				} else {
					if path, _ := exec.LookPath(shellArgs[0]); path != "" {
						fmt.Println(shellArgs[0] + " is " + path)
					} else {
						fmt.Println(shellArgs[0] + ": not found")
					}
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
			if path, _ := exec.LookPath(shellCmd); path != "" {
				cmd := exec.Command(shellCmd, shellArgs...)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
			} else {
				fmt.Println(shellCmd + ": command not found")
			}
		}
	}
}
