package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

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

	var path string
	switch command := os.Args[1]; command {
	case "init":
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
			os.Exit(1)
		}

		fmt.Print("Voidfork repository has been succesfully intialized.")

	case "cat-file":
		if len(os.Args[1:]) < 3 {
			fmt.Fprintf(os.Stderr, "Missing arguments.")
			fmt.Println("usage: voidfork <type> <object>")
			fmt.Println("\t <type>: (-t | -p | -e | -s)")
			fmt.Println("\t <type>: (blob | tree | commit | tag)")
			fmt.Print("\t <object>: blob object to be read")
			os.Exit(1)
		}
		
		cfg := lib.NewCatFileConfig(path)
		err := cfg.CatFile(os.Args[3], os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't read the object. %v", err)
			os.Exit(1)
		}
		
		os.Exit(int(syscall.SIGINT))
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
