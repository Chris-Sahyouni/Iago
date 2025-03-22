package exe

import (
	"iago/src/isa"
)

type Executable interface {
	Info()
	InstructionStream([]segment) []isa.Instruction
}

