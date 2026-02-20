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

			prefixString := string(readBuffer)
			autoCompleteMatches := builtin.AutoComplete(prefixString)

			switch len(autoCompleteMatches) {

			case 0:
				ringBell()

			case 1:
				suffixString := autoCompleteMatches[0][len(prefixString):] + " "
				fmt.Fprint(os.Stdin, suffixString)
				readBuffer = append(readBuffer, suffixString...)
				cursorPtr = len(readBuffer)

			default:
				fmt.Print("\r\n")
				for _, match := range autoCompleteMatches {
					fmt.Printf("%s  ", match)
				}
			}

		case '\r', '\n':
			// '\r' is what I can capture
			// '\n' is what the codecrafters test is capturing
			fmt.Fprint(os.Stdin, "\r\n")
			return string(readBuffer), nil

		default:
			readBuffer = append(readBuffer[:cursorPtr], append(byteBuffer, readBuffer[cursorPtr:]...)...)
			cursorPtr++
			fmt.Fprint(os.Stdin, string(byteBuffer))

		}
	}
}

func redraw(prompt string, readBuffer []byte) {
	fmt.Printf("\r\x1b[K%s %s", prompt, string(readBuffer))
}
