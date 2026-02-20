package main

import (
	"fmt"
	"os"
)

func isRedirectOutput(symbol string) bool {
	return symbol == ">" || symbol == "1>" || symbol == "2>" || symbol == ">>" || symbol == "1>>" || symbol == "2>>"
}

func readFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	return file
}

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	return file
}

func ringBell() {
	fmt.Fprint(os.Stdout, "\x07")
}
