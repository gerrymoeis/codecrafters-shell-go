package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	command = strings.TrimSpace(command)
	if err != nil {
		return
	}
	fmt.Printf("$ %s: command not found\n", command)
}
