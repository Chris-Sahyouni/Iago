package cli

import (
	"errors"
	"iago/src/global"
	"iago/src/term"
)

type CatTarget struct{ args Args }

func (c CatTarget) ValidArgs() bool {
	return len(c.args) == 0
}

func (c CatTarget) Execute(globalState *global.GlobalState) error {

	if globalState.CurrentTarget == struct {
		Title    string
		Contents string
	}{"", ""} {
		return errors.New("no target payload set")
	}

	term.Println("Target Payload:", globalState.CurrentTarget.Title)
	// the contents will need some formatting
	term.Println(globalState.CurrentTarget.Contents)
	return nil
}
