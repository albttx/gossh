package main

import (
	"fmt"
	"os"

	"github.com/albttx/gossh"
)

func main() {
	// Leave password empty for connection with ssh keys
	err := gossh.Prompt("root", "root", "42.42.42.42", "42")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
