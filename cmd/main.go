package main

import (
	"fmt"
	"os"
	"strings"

	"www.github.com/fr13nd230/voidfork/lib"
)

// Usage: voidfork <command> <arg1> <arg2> ...
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Application has been recovered from a panic. %s", r)
			os.Exit(1)
		}
	}()
	
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: voidfork <command> [<args>...]\n")
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "init":
		var path string	
		if len(os.Args) > 2 {		
			path = strings.TrimSpace(os.Args[2])
			if len(path) == 0 {
				path = "./"
			}
		}

		cfg := lib.NewInitConfig(path)
		err := cfg.Init()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't initialize new voidfork repository. %v", err)
		}

		fmt.Print("Voidfork repository has been succesfully intialized.")
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
