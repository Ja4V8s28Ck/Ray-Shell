package builtin

import (
	"fmt"
	"os"
	"strings"
)

type Cd struct{}

func (cd Cd) Name() string {
	return "cd"
}

func (cd Cd) Execute(shellArgs []string) {
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
}
