package main

import (
	"flag"
	"fmt"

	"go-task-example/internal"
)

func main() {
	flag.Parse()
	targets := flag.Args()
	if len(targets) == 0 {
		fmt.Println("Usage: open <http uri>")
		return
	}

	url := flag.Arg(0)
	if err := internal.OpenUrl(url); err != nil {
		fmt.Println(err)
	}
}
