package cli

import (
	"encoding/hex"
	"fmt"
	"iago/src/global"
	"os"
)

type SetPayload struct{ args Args }

func (s SetPayload) ValidArgs() bool {
	if len(s.args) == 0 {
		return true
		// will open interactive editor in this case
	}
	if len(s.args) == 1 {
		_, ok := s.args["default"]
		return ok
	}
	return false
}

func (s SetPayload) Execute(globalState *global.GlobalState) error {
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
	globalState.CurrentPayload.Contents = contentString

	return nil
}