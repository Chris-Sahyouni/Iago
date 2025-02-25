package isa

type ISA interface {
	InstructionSize() int // 1 if instructions are variable length (because you can jump to the middle of the encoding then)
}