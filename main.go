package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mrcruz117/mnky-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Mnky programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
