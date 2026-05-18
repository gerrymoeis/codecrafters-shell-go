package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	command = command[:len(command)-1]
	if err != nil {
		return
	}
	fmt.Printf("%s anjay: command not found\n", command)
}
