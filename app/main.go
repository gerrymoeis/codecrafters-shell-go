package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("there is an error: %v\n", err)
			break
		}
		command = strings.TrimSpace(command)

		if command == "exit" {
			break
		}

		fmt.Printf("%s: command not found\n", command)
	}
}
