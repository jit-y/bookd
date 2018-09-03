package main

import (
	"fmt"
	"os"

	"github.com/jit-y/bookd/cmd"
)

func main() {
	cmd := cmd.NewRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
