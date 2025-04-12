package cli

import (
	"iago/src/global"
	"errors"
	"fmt"
)

type CatPayload struct{ args Args }

func (c CatPayload) ValidArgs() bool {
	return len(c.args) == 0
}

func (c CatPayload) Execute(globalState *global.GlobalState) error {
	if globalState.CurrentPayload == struct {
		Contents string
	}{""} {
		return errors.New("no payload generated yet")
	}
	fmt.Println("Current Payload:")
	fmt.Println(globalState.CurrentPayload.Contents)
	return nil
}