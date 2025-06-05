package main

import (
	"ascii-art/banners"
	"ascii-art/proccesor"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run ./cmd [input string]")
		os.Exit(1)
	}

	errror := banners.CheckASCII(os.Args[1])

	if errror != nil {
		fmt.Println(errror)
		os.Exit(1)
	}

	proccesor.Processing(os.Args[1])
}
