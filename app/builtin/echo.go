package builtin

import (
	"fmt"
	"strings"
)

type Echo struct{}

func (echo Echo) Name() string {
	return "echo"
}

func (echo Echo) Execute(shellArgs []string) {
	fmt.Println(strings.Join(shellArgs, " "))
}
