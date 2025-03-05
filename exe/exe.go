package exe

type Executable interface {
	Info()
	InstructionStream() []Instruction
}

type Instruction struct {
	Op string
	Vaddr uint
}