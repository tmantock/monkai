package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/tmantock/monkai/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! this is the Monkai Programming Language!\n", user.Username)
	fmt.Printf("Please feel free to type in commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}
