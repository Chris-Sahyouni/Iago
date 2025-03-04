package main

import (
	"bytes"
	"errors"
	"fmt"
	"iago/exe"
	"os"
	"bufio"
	"strings"
)

var CurrentFile exe.Executable


func main() {

	fmt.Println("Iago interactive shell")
	fmt.Println("Run help to view available commands")

	reader := bufio.NewReader(os.Stdin)

	for {
		var line string
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			// change this later
			fmt.Println(err)
			continue
		}
		cmd, args, flags := parseUserInput(line)
		executeCommand(cmd, args, flags)
	}

}

func parseUserInput(line string) (string, []string, []string) {
	userInput := strings.TrimSpace(line)
	parsed := strings.Split(userInput, " ")
	cmd := parsed[0]
	var args []string
	var flags []string
	for _, s := range parsed[1:] {
		if len(s) > 2 && s[0:2] == "--" {
			flags = append(flags, s)
		} else {
			args = append(args, s)
		}
	}
	return cmd, args, flags
}

func executeCommand(cmd string, args []string, flags []string) {
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
		fmt.Println("unrecognized command")
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