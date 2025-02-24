package exe

import (
	"iago/isa"
)

type Elf struct {
	instruction_trie isa.ISA
	arch int // either 32 or 64
}

func NewElf(hex_contents string) (*Executable, error) {
	return nil, nil
}
