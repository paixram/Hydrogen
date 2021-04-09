package main

import (
	"fmt"
	"os"

	"github.com/Hydrogen/repl"
)

func main() {
	fmt.Println("Welcome to Hydrogen REPL")
	fmt.Println("Feel free to type whatever you want!!")
	repl.Start(os.Stdin, os.Stdout)
}
