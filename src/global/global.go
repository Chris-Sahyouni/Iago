package global

import (
	"iago/src/exe"
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
	History *History
}
