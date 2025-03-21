package cli

import (
	"fmt"
	"strings"
	"iago/global"
)

type Help struct{ args Args }

func (h Help) ValidArgs() bool {
	return len(h.args) == 0
}

func (Help) Execute(_ *global.GlobalState) error {
	fmt.Println("Commands:")
	fmt.Println("    exit" + strings.Repeat(" ", 16-len("quit")) + "Exit the interactive shell")
	fmt.Println("    help" + strings.Repeat(" ", 16-len("help")) + "Show help")
	fmt.Println("    load <path>" + strings.Repeat(" ", 16-len("load <path>")) + "Sets the current file for analysis")
	fmt.Println("    stat" + strings.Repeat(" ", 16-len("stat")) + "View the current file's metadata")
	fmt.Println("    set-target" + strings.Repeat(" ", 16-len("set-target")) + "")
	fmt.Println("    target" + strings.Repeat(" ", 16-len("target")) + "View the current target payload")
	return nil
}