package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(2)
	}
	cmd := os.Args[1]

	switch cmd {
	case "greet":
		msg := "REMINDER CLI - CLI BASICS"
		if len(os.Args) > 2 {
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
		}
		fmt.Printf("Hello and welcome: %s\n", msg)
	case "help":
		fmt.Println("HELP MESSAGE")
	default:
		fmt.Printf("Unknown command: %s; see 'help' to see all commands\n", cmd)
	}
}
