package global

import (
	"iago/exe"
)

type GlobalState struct {
	CurrentFile   exe.Executable
	TargetPayload struct {
		Title    string
		Contents string
	}
}