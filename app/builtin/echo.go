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
	// fmt.Println(shellArgs)
	fmt.Println(strings.Join(shellArgs, " "))
	// fmt.Println(parseEchoArgs(strings.Join(shellArgs, " ")))
}

func parseEchoArgs(shellArgsString string) string {
	shellArgsString = strings.TrimSpace(shellArgsString)

	stringBuilder := strings.Builder{}
	n := len(shellArgsString)

	isSpace := false
	var quotes rune

	for i := 0; i < n; {
		currChar := rune(shellArgsString[i])

		switch currChar {
		case '\\':
			if quotes == '\'' {
				stringBuilder.WriteRune(currChar)
			}
			stringBuilder.WriteRune(rune(shellArgsString[i+1]))
			i++
		case '"':
		case '\'':
			if quotes == currChar {
				quotes = 0
			} else if quotes == 0 {
				quotes = currChar
			} else {
				stringBuilder.WriteRune(currChar)
			}
		case ' ':
			if quotes != 0 || !isSpace {
				stringBuilder.WriteRune(currChar)
			}
			isSpace = true

		default:
			stringBuilder.WriteRune(currChar)
			isSpace = false
		}

		i++
	}
	return stringBuilder.String()
}
