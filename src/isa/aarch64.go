package isa

type AArch64 struct {}

func (AArch64) InstructionSize() int {
	return 4
}

func (AArch64) GadgetTerminator() string {
	return "d65f03c0"
}

func (AArch64) Name() string {
	return "AArch64 (64-bit ARM)"
}