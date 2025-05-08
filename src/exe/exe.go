package exe

import (
	"github.com/Chris-Sahyouni/iago/src/isa"
	"github.com/Chris-Sahyouni/iago/src/trie"
)

type Executable interface {
	Info()
	InstructionStream([]segment) []isa.Instruction
	ReverseInstructionTrie() *trie.TrieNode
	Endianness() string
	Arch() uint
	Isa() isa.ISA
}

