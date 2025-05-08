package cli

import (
	"errors"
	"github.com/Chris-Sahyouni/iago/src/global"
	"github.com/Chris-Sahyouni/iago/src/term"
	"strings"
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