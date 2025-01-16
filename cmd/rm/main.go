package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	targets := flag.Args()
	if len(targets) == 0 {
		fmt.Println("Usage: rm <file_or_dir>...")
		return
	}

	for _, target := range targets {
		if err := os.RemoveAll(target); err != nil {
			fmt.Printf("Error to remove: %s\n", target)
			return
		}
	}
}
