package builtin

import (
	"os"
)

type Exit struct{}

func (exit Exit) Name() string {
	return "exit"
}

func (exit Exit) Execute(shellArgs []string) {
	os.Exit(0)
}
