package exe

import (
	// "iago/isa"
	// "iago/suffixtree"
	"errors"
)


type Elf struct {
	arch int // either 32 or 64
	endianness string
	// abi string // not sure abi is necessary
}

func (e Elf) foo() {

}

func NewElf(elf []byte) (Executable, error) {
	// determine 32 bit or 64 bit
	var arch int
	var endianness string
	// var abi string

	arch_index := 4
	if (elf[arch_index] == 1) {
		arch = 32
	} else if (elf[arch_index] == 2) {
		arch = 64
	} else {
		return nil, errors.New("invalid ELF file")
	}

	// determine endianness
	endianness_index := 5
	if (elf[endianness_index] == 1) {
		endianness = "little"
	} else if (elf[endianness_index] == 2) {
		endianness = "big"
	} else {
		return nil, errors.New("invalid ELF file")
	}

	// // determine abi
	// abis := []string{"System V", "HP-UX", "NetBSD", "Linux", "GNU Hurd", "Solaris", "AIX", "IRIX", "FreeBSD", "Tru64", "Novell Modesto", "OpenBSD", "OpenVMS", "NonStop Kernel", "AROS", "FenixOS", "Nuxi CloudABI", "Stratus Technologies OpenVOS"}
	// abi_index := 7
	// abi = abis[elf[abi_index]]



	return Elf{
		arch: arch,
		endianness: endianness,
		// abi: abi,
	}, nil
}
