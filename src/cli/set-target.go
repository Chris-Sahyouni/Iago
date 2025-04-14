package cli

import (
	"encoding/hex"
	"iago/src/global"
	"os"
)

type SetTarget struct{ args Args }

func (s SetTarget) ValidArgs() bool {
	// if len(s.args) == 0 {
	// 	// will open interactive editor in this case
	// 	return true
	// }

	if len(s.args) != 1 {
		return false
	}

	_, ok := s.args["default"]
	return ok
}

func (s SetTarget) Execute(globalState *global.GlobalState) error {
	// if len(s.args) == 0 {
	// 	fmt.Println("Would open interactive editor in this case")
	// 	return nil
	// }

	file := s.args["default"]
	contents, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	contentString := hex.EncodeToString(contents)
	globalState.CurrentTarget.Title = file
	globalState.CurrentTarget.Contents = contentString

	return nil
}
