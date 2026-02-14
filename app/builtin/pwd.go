package builtin

import (
	"fmt"
	"os"
)

type Pwd struct{}

func (pwd Pwd) Name() string {
	return "pwd"
}

func (pwd Pwd) Execute(shellArgs []string) {
	wd, err := os.Getwd()

	if err != nil {
		fmt.Printf("pwd: %v", wd)
	} else {
		fmt.Println(wd)
	}
}
