package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print(command, err)
	fmt.Printf("$ %s: command not found\n", command)
}
