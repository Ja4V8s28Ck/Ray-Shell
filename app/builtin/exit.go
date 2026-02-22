package builtin

import (
	"os"
)

type Exit struct{}

func (exit Exit) Name() string {
	return "exit"
}

func (exit Exit) Execute(shellArgs []string, ctx *ExecContext) {
	// Pre-exit history append/write
	if HISTFILE, ok := os.LookupEnv("HISTFILE"); ok {

		_, err := os.Stat(HISTFILE)
		if os.IsExist(err) {
			AppendHistoryToFile(HISTFILE)
		} else {
			WriteHistoryToFile(HISTFILE)
		}
	}

	os.Exit(0)
}
