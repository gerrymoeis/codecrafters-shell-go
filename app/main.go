package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(buildFullPath())

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
				fmt.Printf("%s: not found\n", args)
			}
			// else if  {
			// 	fmt.Printf("%s is %s\n", args, full_path)
			// }

		} else {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}

func buildFullPath() map[string]string {
	path := os.Getenv("PATH")
	full_path := map[string]string{}
	for _, dir := range strings.Split(path, ";") {
		fmt.Println(dir)
		full_path[dir] = dir
	}
	return full_path
}
