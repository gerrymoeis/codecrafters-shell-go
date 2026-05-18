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

		command, args, found := strings.Cut(input, " ")
		if !found {
			command = input
			args = ""
		}

		if command == "exit" {
			break
		} else if command == "echo" {
			fmt.Println(args)
		} else {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
