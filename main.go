package main

import (
	"bytes"
	"errors"
	"fmt"
	"iago/exe"
	"os"
	"strings"
)

var CurrentFile exe.Executable


func main() {

	fmt.Println("Iago interactive shell")
	fmt.Println("Run help to view available commands")


	for {
		var userInput string
		fmt.Print("> ")
		fmt.Scanln(&userInput)
		parseUserInput(userInput)
	}

}

func parseUserInput(userInput string) {
	parsed := strings.Split(userInput, " ")
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
	default:
		fmt.Println("Unrecognized command")
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
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	exeConstructors := map[string]func([]byte) (exe.Executable, error){
		"elf": exe.NewElf,
	}
	fileType, err := DetermineFileType(fileBytes)
	if err != nil {
		return err
	}
	constructor := exeConstructors[fileType]
	newExecutable, err := constructor(fileBytes)
	if err != nil {
		return err
	}
	CurrentFile = newExecutable

	return nil
}

func DetermineFileType(fileBytes []byte) (string, error) {
	elfMagic := []byte{'\x7f', '\x45', '\x4c', '\x46'}
	if bytes.Equal(fileBytes[:4], elfMagic) {
		return "elf", nil
	}

	return "", errors.New("unrecognized file format")
}


func Help() {
	fmt.Println("Commands:")
	fmt.Println("    exit" + strings.Repeat(" ", 16 - len("quit")) + "Exit the interactive shell")
	fmt.Println("    help" + strings.Repeat(" ", 16 - len("help")) + "Show help")
	fmt.Println("    load <path>" + strings.Repeat(" ", 16 - len("load <path>")) + "Sets the current file for analysis")
}