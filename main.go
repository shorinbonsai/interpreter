package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s! This is an interpreter.\n", user.Username)
	fmt.Printf("Do the typy typy stuff here.\n")
	repl.Start(os.Stdin, os.Stdout)
}
