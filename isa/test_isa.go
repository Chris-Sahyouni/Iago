package isa

type TestISA struct {}

func (TestISA) InstructionSize() int {
	return 1
}

func (TestISA) GadgetTerminator() string {
	return "z"
}

func (TestISA) Name() string {
	return "test"
}