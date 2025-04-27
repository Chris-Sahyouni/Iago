package cli

import (
	"errors"
	"fmt"
	"iago/src/global"
	"iago/src/term"
)

type CatPayload struct{ args Args }

func (c CatPayload) ValidArgs() bool {
	return len(c.args) == 0
}

func (c CatPayload) Execute(globalState *global.GlobalState) error {
	// if (globalState.CurrentPayload.Chain) == struct {
	// 	PaddingLength int
	// 	Chain []uint
	// }{0, nil} {
	// 	return errors.New("no payload generated yet")
	// }

	if globalState.CurrentPayload.Chain == nil {
		return errors.New("no payload generated or set")
	}

	term.Println("Current Payload:")
	term.Println("  Padding Bytes:", globalState.CurrentPayload.PaddingLength)
	term.Println("  Chain:")
	for _, gaddr := range globalState.CurrentPayload.Chain {
		fmt.Printf("    %x\n", gaddr)
	}
	return nil
}
