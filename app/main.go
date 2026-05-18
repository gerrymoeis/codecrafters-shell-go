package main

import (
	"fmt"
)

func main() {
	var command string
	fmt.Scan(&command)
	fmt.Printf("$%s: command not found", command)
}
