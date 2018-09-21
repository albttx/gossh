package main

import (
	"fmt"
	"os"

	"github.com/albttx/gossh"
)

var (
	user string = "root"
	pass string = "root"
	host string = ""
	port string = "22"
)

func main() {
	// Leave password empty for connection with ssh keys
	err := gossh.Exec(user, pass, host, port, "echo 'Hello World !'")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = gossh.Prompt(user, pass, host, port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
