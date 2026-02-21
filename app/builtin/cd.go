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

func (cd Cd) Execute(shellArgs []string, ctx *ExecContext) {
	shellArgsCount := len(shellArgs)

	if shellArgsCount > 1 { // Only 1 arg allowed
		fmt.Fprintln(ctx.Stderr, "cd: too many arguments")
		return
	}

	if shellArgsCount == 0 {
		shellArgs = append(shellArgs, "/")
	}

	cdArg := shellArgs[0]

	if cdArg == "~" || strings.HasPrefix(cdArg, "~/") { // convert relative path to abs path
		cdArg = strings.Replace(cdArg, "~", os.Getenv("HOME"), 1)
	}

	if err := os.Chdir(cdArg); err != nil {
		fmt.Fprintf(ctx.Stderr, "cd: %s: No such file or directory\n", cdArg)
	}
}
