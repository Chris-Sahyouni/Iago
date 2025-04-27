package main

import (
	"fmt"
	"iago/src/cli"
	"iago/src/global"
	"iago/src/term"
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
			Chain         []uint
		}{
			PaddingLength: 0,
			Chain:         nil,
		},
		History: &history,
	}

	terminal := term.RawTerminal()

	for {
		line, err := terminal.ReadLine()
		cmd, err := cli.ParseLine(line)
		if err != nil {
			term.Println(err)
			continue
		}
		if !cmd.ValidArgs() {
			term.Println("invalid arguments")
		}
		err = cmd.Execute(&globalState)
		if err != nil {
			term.Println(err)
		}

	}

}
