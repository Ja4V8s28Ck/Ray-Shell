package builtin

import (
	"fmt"
	"os"
)

type Pwd struct{}

func (pwd Pwd) Name() string {
	return "pwd"
}

func (pwd Pwd) Execute(shellArgs []string, ctx *ExecContext) {
	wd, err := os.Getwd()

	if err != nil {
		fmt.Fprintf(ctx.Stderr, "pwd: %v\n", err)
	} else {
		fmt.Fprintln(ctx.Stdout, wd)
	}
}
