package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buffReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		shellCmd, shellArgs := readCmd(buffReader)

		execCmd(shellCmd, shellArgs)
	}
}
