package utils

import (
	"fmt"
	"os"
)

func RingBell() {
	fmt.Fprint(os.Stdout, "\x07")
}

func IsRedirectOutput(symbol string) bool {
	return symbol == ">" || symbol == "1>" || symbol == "2>" || symbol == ">>" || symbol == "1>>" || symbol == "2>>"
}

func ReadFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	return file
}

func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	return file
}

func FindLongestCommonPrefix(autoCompleteMatches []string) string {
	referenceString := autoCompleteMatches[0]

	for i := range referenceString {

		for _, match := range autoCompleteMatches[1:] {
			if match[i] != referenceString[i] {
				return referenceString[:i]
			}
		}
	}
	return referenceString
}
