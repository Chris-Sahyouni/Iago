package main

import (
	"fmt"
	"strings"
	"os"
	"errors"
	"iago/exe"
	"encoding/hex"
)

var current_file *exe.Executable


func main() {

	fmt.Println("Iago interactive shell")
	fmt.Println("Run help to view available commands")


	for {
		var user_input string
		fmt.Print("> ")
		fmt.Scanln(&user_input)
		parse_user_input(user_input)
	}

}

func parse_user_input(user_input string) {
	parsed := strings.Split(user_input, " ")
	cmd := parsed[0]
	var args []string
	if len(parsed) > 1 {
		args = parsed[1:]
	}
	switch cmd {
	case "help":
		Help()
	case "quit":
		os.Exit(0)
	case "load":
		err := Load(args)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}

func Load(args []string) error {
	if len(args) == 0 {
		return errors.New("missing 1 argument: <path>")
	}
	if len(args) > 1 {
		return errors.New("too many arguments given")
	}

	path := args[0]
	file_bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	hex_contents := hex.EncodeToString(file_bytes)
	exe_constructors := map[string]func(string) (*exe.Executable, error){
		"elf": exe.NewElf,
	}
	file_type, err := DetermineFileType(hex_contents)
	if err != nil {
		return err
	}
	constructor := exe_constructors[file_type]
	new_executable_ptr, err := constructor(hex_contents)
	if err != nil {
		return err
	}
	current_file = new_executable_ptr

	return nil
}

// determine the file type and parse it into the relevant struct
func DetermineFileType(hex_contents string) (string, error) {
	elf_magic := "7f454c46"
	if hex_contents[:4] == elf_magic {
		return "elf", nil
	}

	return "", errors.New("unrecognized file format")
}

func Help() {
	fmt.Println("Commands:")
	fmt.Println("    help" + strings.Repeat(" ", 16 - len("help")) + "Show help")
	fmt.Println("    exit" + strings.Repeat(" ", 16 - len("quit")) + "Exit the interactive shell")
	fmt.Println("    load <path>" + strings.Repeat(" ", 16 - len("load <path>")) + "Sets the current file for analysis")
}