package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/builtin"
	"golang.org/x/term"
)

func readLine(prompt string) (string, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error initalizing xterm")
		os.Exit(1)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Print(prompt + " ")
	var readBuffer []byte

	cursorPtr := 0

	for {
		byteBuffer := make([]byte, 1)
		os.Stdin.Read(byteBuffer)

		switch byteBuffer[0] {

		case 0x03:
			return "", fmt.Errorf("^C")

		case 0x7f:
			if cursorPtr > 0 {
				readBuffer = append(readBuffer[:cursorPtr-1], readBuffer[cursorPtr:]...)
				cursorPtr--
				redraw(prompt, readBuffer)
			}

		case '\t':
			if len(readBuffer) == 0 { // don't autocomplete when the buffer is empty
				continue
			}

			autoCompleteMatches := builtin.AutoComplete(string(readBuffer))
			if len(autoCompleteMatches) == 1 {
				fmt.Print(autoCompleteMatches[0] + " ")
			}

		case '\r':
			fmt.Print("\r\n")
			return string(readBuffer), nil

		default:
			readBuffer = append(readBuffer[:cursorPtr], append(byteBuffer, readBuffer[cursorPtr:]...)...)
			cursorPtr++
			fmt.Print(string(byteBuffer))

		}
	}
}

func redraw(prompt string, readBuffer []byte) {
	fmt.Printf("\r\x1b[K%s %s", prompt, string(readBuffer))
}
