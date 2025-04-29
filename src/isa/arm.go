package isa

type ARM struct {}

func (ARM) InstructionSize() int {
	return 4
}

func (ARM) GadgetTerminator() string {
	return "4770"
}

func (ARM) Name() string {
	return "ARM"
}