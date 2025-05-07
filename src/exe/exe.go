package exe

import (
	"iago/src/isa"
	"iago/src/trie"
)

type Executable interface {
	Info()
	InstructionStream([]segment) []isa.Instruction
	ReverseInstructionTrie() *trie.TrieNode
	Endianness() string
	Arch() uint
	Isa() isa.ISA
}

