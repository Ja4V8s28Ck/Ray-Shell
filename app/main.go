package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	buffReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command, err := buffReader.ReadString('\n')

		shellCommand := command[:len(command)-1]
		if shellCommand == "exit" {
			os.Exit(0)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		fmt.Println(shellCommand + ": command not found")
	}
}
