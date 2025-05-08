package exe

import (
	"github.com/Chris-Sahyouni/iago/isa"
	"github.com/Chris-Sahyouni/iago/trie"
)

type Executable interface {
	Info()
	InstructionStream([]segment) []isa.Instruction
	ReverseInstructionTrie() *trie.TrieNode
	Endianness() string
	Arch() uint
	Isa() isa.ISA
}
