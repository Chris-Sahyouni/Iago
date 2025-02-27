package isa

type x86 struct {}

func (x86) InstructionSize() int {
	return 1
}

func (x86) GadgetTerminator() string {
	return "c3"
}

