package builtin

import (
	"fmt"
	"strconv"
)

type History struct{}

func (history History) Name() string {
	return "history"
}

var historyArr []string

func (history History) Execute(shellArgs []string, ctx *ExecContext) {
	shellArgsCount := len(shellArgs)

	limit := 0

	if shellArgsCount > 1 {
		fmt.Fprintln(ctx.Stderr, "history: too many arguments")
		return
	}

	if shellArgsCount == 1 {
		if intVal, err := strconv.Atoi(shellArgs[0]); err == nil {
			limit = intVal
		} else {
			fmt.Fprintln(ctx.Stderr, "history: invalid argument")
		}
	}

	historyArrCount := len(historyArr)

	if limit == 0 || limit > historyArrCount {
		limit = historyArrCount
	}

	for idx, historyCmd := range historyArr[historyArrCount-limit:] {
		fmt.Fprintf(ctx.Stdout, "    %d  %s\n", idx+1+(historyArrCount-limit), historyCmd)
	}
}

func StoreHistory(cmdLine string) {
	historyArr = append(historyArr, cmdLine)
}
