package cli

import (
	"errors"
	"github.com/Chris-Sahyouni/iago/src/global"
	"github.com/Chris-Sahyouni/iago/src/term"
	"strings"
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

func (CatTarget) Help() {
	term.Println("    cat-target" + strings.Repeat(" ", SPACE_BETWEEN-len("cat-target")) + "View the current target payload")
}

