package cli

import (
	"fmt"
	"iago/src/global"
	"strings"
)

const SPACE_BETWEEN = 24

type Help struct{ args Args }

func (h Help) ValidArgs() bool {
	return len(h.args) == 0
}

func (Help) Execute(_ *global.GlobalState) error {
	fmt.Println("Commands:")
	fmt.Println("    exit" + strings.Repeat(" ", SPACE_BETWEEN-len("quit")) + "Exit the interactive shell")
	fmt.Println("    help" + strings.Repeat(" ", SPACE_BETWEEN-len("help")) + "Show help")
	fmt.Println("    load <path>" + strings.Repeat(" ", SPACE_BETWEEN-len("load <path>")) + "Sets the current file for analysis")
	fmt.Println("    rop [OPTIONS]" + strings.Repeat(" ", SPACE_BETWEEN-len("rop [OPTIONS]")) + "Generate a ROP chain for the target file and payload")
	fmt.Println("        -o FILE  Specify output file (optional)")
	fmt.Println("    stat" + strings.Repeat(" ", SPACE_BETWEEN-len("stat")) + "View the current file's metadata")
	fmt.Println("    cat-target" + strings.Repeat(" ", SPACE_BETWEEN-len("cat-target")) + "View the current target payload")
	fmt.Println("    set-target <path>" + strings.Repeat(" ", SPACE_BETWEEN-len("set-target <path>")) + "Set the target payload")
	fmt.Println("    cat-payload" + strings.Repeat(" ", SPACE_BETWEEN-len("cat-payload")) + "View the current payload")
	fmt.Println("    set-payload <path>" + strings.Repeat(" ", SPACE_BETWEEN-len("set-payload <path>")) + "Set the current payload")
	return nil
}
