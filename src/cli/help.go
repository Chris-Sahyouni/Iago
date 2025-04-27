package cli

import (
	"iago/src/global"
	"iago/src/term"
	"strings"
)

const SPACE_BETWEEN = 24

type Help struct{ args Args }

func (h Help) ValidArgs() bool {
	return len(h.args) == 0
}

func (Help) Execute(_ *global.GlobalState) error {
	term.Println("Commands:")
	term.Println("    help" + strings.Repeat(" ", SPACE_BETWEEN-len("help")) + "View this help menu")
	term.Println("    load <path>" + strings.Repeat(" ", SPACE_BETWEEN-len("load <path>")) + "Sets the current file for analysis")
	term.Println("    stat" + strings.Repeat(" ", SPACE_BETWEEN-len("stat")) + "View the current file's metadata")
	term.Println("    cat-target" + strings.Repeat(" ", SPACE_BETWEEN-len("cat-target")) + "View the current target payload")
	term.Println("    set-target <path>" + strings.Repeat(" ", SPACE_BETWEEN-len("set-target <path>")) + "Set the target payload")
	term.Println("    cat-payload" + strings.Repeat(" ", SPACE_BETWEEN-len("cat-payload")) + "View the current payload")
	term.Println("    set-payload <path>" + strings.Repeat(" ", SPACE_BETWEEN-len("set-payload <path>")) + "Set the current payload. If the desired payload has been prepended with padding, the file name must contain the substring \"pad\" followed by the number of bytes of padding. e.g. \"pad64\"")
	term.Println("    rop [OPTIONS]" + strings.Repeat(" ", SPACE_BETWEEN-len("rop [OPTIONS]")) + "Generate a ROP chain for the target file and payload")
	term.Println("        -o FILE  Specify output file (optional)")
	term.Println("    pad <bytes>" + strings.Repeat(" ", SPACE_BETWEEN-len("pad <bytes>")) + "Generate a new payload by prepending <bytes> number of bytes of padding to the current payload.")
	term.Println("    quit" + strings.Repeat(" ", SPACE_BETWEEN-len("quit")) + "Exit the interactive shell")

	return nil
}
