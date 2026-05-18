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

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("there is an error: %v\n", err)
			break
		}
		input = strings.TrimSpace(input)

		types := map[string]struct{}{
			"exit": {},
			"echo": {},
			"type": {},
		}

		command, args, _ := strings.Cut(input, " ")
		if command == "exit" {
			break
		} else if command == "echo" {
			fmt.Println(args)
		} else if command == "type" {
			if _, exists := types[args]; exists {
				fmt.Printf("%s is a shell builtin\n", args)
			} else {
				errorMsg(args)
			}
		} else {
			errorMsg(command)
		}
	}
}

func errorMsg(command string) {
	fmt.Printf("%s: command not found\n", command)
}
