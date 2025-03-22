package global

import (
	"iago/src/exe"
)

type GlobalState struct {
	CurrentFile   exe.Executable
	TargetPayload struct {
		Title    string
		Contents string
	}
}