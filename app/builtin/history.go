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

var HistoryArrCount int

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
			return
		}
	}

	if limit == 0 || limit > HistoryArrCount {
		limit = HistoryArrCount
	}

	for idx, historyCmd := range historyArr[HistoryArrCount-limit:] {
		fmt.Fprintf(ctx.Stdout, "    %d  %s\n", idx+1+(HistoryArrCount-limit), historyCmd)
	}
}

func StoreHistory(cmdLine string) {
	historyArr = append(historyArr, cmdLine)
	HistoryArrCount++
}

func GetHistory(historyArrIdx *int, direction rune) string {

	var historyString string
	if direction == 'u' {
		*historyArrIdx--
		historyString = historyArr[*historyArrIdx]
	} else {
		*historyArrIdx++
		historyString = historyArr[*historyArrIdx]
	}

	return historyString
}
