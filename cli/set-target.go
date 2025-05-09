package cli

import (
	"os"
	"strings"

	"github.com/Chris-Sahyouni/iago/global"
	"github.com/Chris-Sahyouni/iago/term"
)

type SetTarget struct{ args Args }

func (s SetTarget) ValidArgs() bool {

	if len(s.args) == 0 {
		return true
	}

	if len(s.args) != 1 {
		return false
	}

	_, ok := s.args["default"]
	return ok
}

func (s SetTarget) Execute(globalState *global.GlobalState) error {

	if len(s.args) == 0 {
		term.Println("input target:")
		target, err := globalState.Terminal.ReadLine()
		if err != nil {
			return err
		}
		globalState.CurrentTarget.Contents = target
		globalState.CurrentTarget.Title = ""
		term.Println("target set to: " + target)
		return nil
	}

	file := s.args["default"]
	contents, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	contentString := string(contents)

	term.Println(contentString)

	globalState.CurrentTarget.Title = file
	globalState.CurrentTarget.Contents = contentString

	return nil
}

func (SetTarget) Help() {
	term.Println("    set-target <path>" + strings.Repeat(" ", SPACE_BETWEEN-len("set-target <path>")) + "Set the target payload. Alternatively, exclude <path> to manually input the target")
}
