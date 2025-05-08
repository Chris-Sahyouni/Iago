package cli

import (
	"errors"
	"strings"

	"github.com/Chris-Sahyouni/iago/global"
	"github.com/Chris-Sahyouni/iago/term"
)

type Stat struct{ args Args }

func (s Stat) ValidArgs() bool {
	return len(s.args) == 0
}

func (Stat) Execute(globalState *global.GlobalState) error {
	if globalState.CurrentFile == nil {
		return errors.New("no file loaded")
	}
	globalState.CurrentFile.Info()
	return nil
}

func (Stat) Help() {
	term.Println("    stat" + strings.Repeat(" ", SPACE_BETWEEN-len("stat")) + "View the current file's metadata")
}
