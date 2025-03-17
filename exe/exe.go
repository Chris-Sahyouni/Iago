package exe

import (
	"iago/isa"
)

type Executable interface {
	Info()
	InstructionStream() []isa.Instruction
}

