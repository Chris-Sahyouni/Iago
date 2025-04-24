package isa

type TestISA struct {}

func (TestISA) InstructionSize() int {
	return 1
}

func (TestISA) GadgetTerminator() string {
	return "z_"
}

func (TestISA) Name() string {
	return "test"
}