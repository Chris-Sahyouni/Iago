package isa

type ISA interface {
	InstructionSize() int // 1 if instructions are of variable length (because you can execute from the middle of an encoding in this case)
	GadgetTerminator() string
	Name() string
}

type Instruction struct {
	Op string
	Vaddr uint
}