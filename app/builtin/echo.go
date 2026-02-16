package builtin

import (
	"fmt"
	"strings"
)

type Echo struct{}

func (echo Echo) Name() string {
	return "echo"
}

func (echo Echo) Execute(shellArgs []string, ctx *ExecContext) {
	echoOut := strings.Join(shellArgs, " ")
	fmt.Fprintln(ctx.Stdout, echoOut)
}
