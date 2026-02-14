package builtin

import (
	"fmt"
	"os/exec"
)

type Type struct{}

func (typeCmd Type) Name() string {
	return "type"
}

func (typeCmd Type) Execute(shellArgs []string) {
	if len(shellArgs) == 0 {
		fmt.Println("type: needs an argument")

	} else if len(shellArgs) == 1 {

		if _, ok := CmdFuncMap[shellArgs[0]]; ok {
			fmt.Println(shellArgs[0] + " is a shell builtin")

		} else {
			if path, _ := exec.LookPath(shellArgs[0]); path != "" {
				fmt.Println(shellArgs[0] + " is " + path)
			} else {
				fmt.Println(shellArgs[0] + ": not found")
			}
		}
	}
}
