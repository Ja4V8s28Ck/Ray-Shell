package builtin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/codecrafters-io/shell-starter-go/app/utils"
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

	if shellArgsCount == 2 && len(shellArgs[0]) == 2 && shellArgs[0][0] == '-' {
		// handle flags
		switch shellArgs[0][1] {

		case 'r':
			ReadHistoryFromFile(shellArgs[1])

		case 'w':
			WriteHistoryToFile(shellArgs[1])

		default:
			fmt.Fprintf(ctx.Stderr, "history: %s: invalid option\n", shellArgs[0])

		}
		return
	}

	// if shellArgsCount > 1 { // add as else condition and kick out
	// 	fmt.Fprintln(ctx.Stderr, "history: too many arguments")
	// 	return
	// }

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
	if cmdLine == "" {
		return
	}

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

func ReadHistoryFromFile(fileName string) {
	file := utils.ReadFile(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		StoreHistory(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "history: %v \n", err)
	}
}

func WriteHistoryToFile(fileName string) {
	file := utils.CreateFile(fileName)
	defer file.Close()

	scanner := bufio.NewWriter(file)
	for _, history := range historyArr {
		if _, err := scanner.WriteString(history + "\n"); err != nil {
			fmt.Fprintf(os.Stderr, "history: %v \n", err)
		}
	}

	if err := scanner.Flush(); err != nil {
		fmt.Fprintf(os.Stderr, "history: %v \n", err)
	}
}
