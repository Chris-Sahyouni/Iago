package isa

type Thumb struct {}

func (Thumb) InstructionSize() int {
	return 2
}

func (Thumb) GadgetTerminator() string {
	return "4770"
}

func (Thumb) Name() string {
	return "ARM (Thumb mode)"
}