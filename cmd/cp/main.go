package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	targets := flag.Args()
	if len(targets) != 2 {
		fmt.Println("Usage: rm <source> <destination>")
		return
	}

	source := flag.Arg(0)
	destination := flag.Arg(1)
	sourceInfo, err := os.Stat(source)
	if err != nil {
		fmt.Printf("Error to copy %s: %v\n", source, err)
		return
	}

	if sourceInfo.IsDir() {
		err = copyDir(source, destination)
	} else {
		err = copyFile(source, destination)
	}

	if err != nil {
		fmt.Println(err)
	}
}

func copyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}
	err = os.Chmod(destination, sourceFileInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}

func copyDir(source, destination string) error {
	err := os.MkdirAll(destination, 0644)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(source, entry.Name())
		destinationPath := filepath.Join(destination, entry.Name())

		if entry.IsDir() {
			err = copyDir(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
