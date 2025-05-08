package global

import (
	"github.com/Chris-Sahyouni/iago/src/exe"
	"github.com/Chris-Sahyouni/iago/src/term"
)

type GlobalState struct {
	CurrentFile   exe.Executable
	CurrentTarget struct {
		Title    string
		Contents string
	}
	CurrentPayload struct {
		PaddingLength int
		Chain         []uint
	}
	Terminal term.Terminal
}
