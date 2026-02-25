package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/builtin"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
	"golang.org/x/term"
)

var prompt = "$"

func readLine() (string, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error initalizing xterm")
		os.Exit(1)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Print(prompt + " ")
	var readBuffer []byte
	var tmpReadBuffer []byte
	var isArrowKeyPressed bool

	cursorPtr := 0
	tabCount := 0

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
				redraw(readBuffer)

				tabCount = 0 // reset tabcount
			}

		case '\t':
			if len(readBuffer) == 0 { // don't autocomplete when the buffer is empty
				continue
			}

			prefixString := string(readBuffer)
			var autoCompleteMatches []string

			if spaceIdx := strings.LastIndex(prefixString, " "); spaceIdx != -1 {
				if spaceIdx+1 < len(prefixString) {
					prefixString = prefixString[spaceIdx+1:]
				} else {
					prefixString = ""
				}
				autoCompleteMatches = builtin.FileAutoComplete(prefixString)

			} else {
				autoCompleteMatches = builtin.AutoComplete(prefixString)
			}

			prefixStringLen := len(prefixString)

			totalMatches := len(autoCompleteMatches)

			switch totalMatches {
			case 0:
				utils.RingBell()

			case 1:
				suffixString := autoCompleteMatches[0][prefixStringLen:] + " "
				fmt.Fprint(os.Stdin, suffixString)
				readBuffer = append(readBuffer, suffixString...)
				cursorPtr = len(readBuffer)

			default:
				// add longest common prefix if there is one
				longestCommonPrefix := utils.FindLongestCommonPrefix(autoCompleteMatches)
				suffixString := longestCommonPrefix[prefixStringLen:]
				if suffixString != "" {
					fmt.Fprint(os.Stdin, suffixString)
					readBuffer = append(readBuffer, suffixString...)
					cursorPtr = len(readBuffer)

				} else if tabCount == 0 {
					utils.RingBell()
					tabCount++

				} else {
					fmt.Fprintf(os.Stdout, "\r\n%v\n", strings.Join(autoCompleteMatches, "  "))
					redraw(readBuffer)
				}
			}

		case '\r', '\n':
			// '\r' is what I can capture
			// '\n' is what the codecrafters test is capturing
			fmt.Fprint(os.Stdin, "\r\n")
			return string(readBuffer), nil

		case 0x1b:
			seq := make([]byte, 2)
			os.Stdin.Read(seq)
			if seq[0] == '[' {

				isArrowKeyPressed = true // track arrowkeys to handle temp buffer

				switch seq[1] {

				case 'A':
					if 0 >= historyArrPtr {
						continue
					}
					historyBuffer := builtin.GetHistory(&historyArrPtr, 'u')
					readBuffer = []byte(historyBuffer)
					redraw(readBuffer)
					cursorPtr = len(readBuffer)

				case 'B':
					if historyArrPtr >= builtin.HistoryArrCount-1 {
						historyArrPtr = builtin.HistoryArrCount // index correction
						readBuffer = tmpReadBuffer
						redraw(readBuffer)
						cursorPtr = len(readBuffer)
						continue
					}

					historyBuffer := builtin.GetHistory(&historyArrPtr, 'd')
					readBuffer = []byte(historyBuffer)
					redraw(readBuffer)
					cursorPtr = len(readBuffer)

				}
			}

		default:
			readBuffer = append(readBuffer[:cursorPtr], append(byteBuffer, readBuffer[cursorPtr:]...)...)
			cursorPtr++
			fmt.Fprint(os.Stdin, string(byteBuffer))

			tabCount = 0 // reset tabCount
		}

		if !isArrowKeyPressed {
			tmpReadBuffer = readBuffer
		}
		isArrowKeyPressed = false
	}
}

func redraw(readBuffer []byte) {
	fmt.Fprintf(os.Stdin, "\r\x1b[K%s %s", prompt, string(readBuffer))
}
