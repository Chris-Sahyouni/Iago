package isa

type ARM struct {}

func (ARM) InstructionSize() int {
	return 4
}

func (ARM) GadgetTerminator() string {
	return "e12fff1e"
}

func (ARM) Name() string {
	return "ARM"
}