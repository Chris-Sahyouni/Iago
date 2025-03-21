package cli

import (
	"os"
	"iago/global"
)

type Quit struct{ args Args }

func (q Quit) ValidArgs() bool {
	return len(q.args) == 0
}

func (Quit) Execute(_ *global.GlobalState) error {
	os.Exit(0)
	return nil
}


