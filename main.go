package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {

	fmt.Println("Iago interactive shell")
	fmt.Println("Run help to view available commands")


	for {
		var user_input string
		fmt.Print("> ")
		fmt.Scanln(&user_input)
		parse(user_input)
	}

}

func parse(user_input string) {
	parsed := strings.Split(user_input, " ")
	cmd := parsed[0]
	// var args []string
	// if len(parsed) > 1 {
	// 	args = parsed[1:]
	// }
	switch cmd {
	case "help":
		help()
	case "quit":
		os.Exit(0)
	}
}

func help() {
	fmt.Println("Commands:")
	fmt.Println("    help" + strings.Repeat(" ", 16 - len("help")) + "Show help")
	fmt.Println("    exit" + strings.Repeat("", 16 - len("quit")) + "Exit the interactive shell")
}