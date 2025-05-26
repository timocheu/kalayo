package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/timocheu/kalayo/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hallo %s, Play with the kalayo\n", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
