package cli

import (
	"errors"
	"iago/global"
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
