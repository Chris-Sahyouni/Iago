package exe

type Executable interface {
	Info()
	InstructionStream() []Instruction
}

type Instruction struct {
	Op string
	Vaddr uint
}

// func (i Instruction) Equals(o Instruction) bool {
// 	return i.Op == o.Op
// }