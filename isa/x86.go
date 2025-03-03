package isa

type X86 struct {}

func (X86) InstructionSize() int {
	return 1
}

func (X86) GadgetTerminator() string {
	return "c3"
}

