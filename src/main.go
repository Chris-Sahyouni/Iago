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

	history := global.History{}
	history.Init()

	globalState := global.GlobalState{
		CurrentFile: nil,
		CurrentTarget: struct {
			Title    string
			Contents string
		}{
			Title:    "",
			Contents: "",
		},
		CurrentPayload: struct {
			PaddingLength int
			Chain []uint
		}{
			PaddingLength: 0,
			Chain: nil,
		},
		History: &history,
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		var line string
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		history.Add(line)

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