package isa

type x86 struct {}

func (_ x86) InstructionSize() int {
	return 1
}
