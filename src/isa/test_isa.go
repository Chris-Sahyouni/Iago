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

type TestISA2 struct {}

func (TestISA2) InstructionSize() int {
	return 2
}

func (TestISA2) GadgetTerminator() string {
	return "zz__"
}

func (TestISA2) Name() string {
	return "test2"
}