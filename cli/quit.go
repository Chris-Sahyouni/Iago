package cli

import (
	"os"
	"strings"

	"github.com/Chris-Sahyouni/iago/global"
	"github.com/Chris-Sahyouni/iago/term"
)

type Quit struct{ args Args }

func (q Quit) ValidArgs() bool {
	return len(q.args) == 0
}

func (Quit) Execute(globalState *global.GlobalState) error {
	globalState.Terminal.Restore()
	os.Exit(0)
	return nil
}

func (Quit) Help() {
	term.Println("    quit" + strings.Repeat(" ", SPACE_BETWEEN-len("quit")) + "Exit the interactive shell")
}
