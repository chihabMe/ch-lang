package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/chihabMe/mg-lang/repel"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s you are using mg-lang \n ", user.Username)
	repel.Start(os.Stdin, os.Stdout)
}
