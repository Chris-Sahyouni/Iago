package cli

import (
	"encoding/hex"
	"fmt"
	"iago/src/global"
	"os"
)

type SetTarget struct{ args Args }

func (s SetTarget) ValidArgs() bool {
	if len(s.args) == 0 {
		// will open interactive editor in this case
		return true
	}
	if len(s.args) == 1 {
		_, ok := s.args["default"]
		return ok
	}
	return false
}

func (s SetTarget) Execute(globalState *global.GlobalState) error {
	if len(s.args) == 0 {
		fmt.Println("Would open interactive editor in this case")
		return nil
	}

	file := s.args["default"]
	contents, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	contentString := hex.EncodeToString(contents)
	globalState.TargetPayload.Title = file
	globalState.TargetPayload.Contents = contentString

	return nil
}
