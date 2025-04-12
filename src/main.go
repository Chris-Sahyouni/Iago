package main

import (
	"bufio"
	"fmt"
	"iago/src/cli"
	"iago/src/global"
	"os"
)

func main() {

	fmt.Println("Iago interactive shell")
	fmt.Println("Run help to view available commands")

	reader := bufio.NewReader(os.Stdin)
	globalState := global.GlobalState{
		CurrentFile: nil,
		CurrentTarget: struct {
			Title    string
			Contents string
		}{
			Title:    "",
			Contents: "",
		},
	}

	for {
		var line string
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			// change this later
			fmt.Println(err)
			continue
		}
		cmd, err := cli.ParseLine(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !cmd.ValidArgs() {
			fmt.Println("invalid arguments")
		}
		err = cmd.Execute(&globalState)
		if err != nil {
			fmt.Println(err)
		}
	}

}
