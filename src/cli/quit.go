package cli

import (
	"os"
	"iago/src/global"
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


