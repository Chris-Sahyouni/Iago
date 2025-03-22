package cli

import (
	"errors"
	"fmt"
	"iago/src/global"
)

type CatTarget struct {args Args}

func (c CatTarget) ValidArgs() bool {
	return len(c.args) == 0
}

func (c CatTarget) Execute(globalState *global.GlobalState) error {

	if globalState.TargetPayload == struct{Title string; Contents string}{"", ""} {
		return errors.New("no target payload set")
	}

	fmt.Println("Target Payload:", globalState.TargetPayload.Title)
	// the contents will need some formatting
	fmt.Println(globalState.TargetPayload.Contents)
	return nil
}