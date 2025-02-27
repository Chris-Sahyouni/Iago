package exe

import (
	// "iago/isa"
	// "iago/suffixtree"
	"encoding/binary"
	"errors"
)

type Elf struct {
	arch       uint // either 32 or 64
	endianness string
}

func (e Elf) foo() {

}

func NewElf(elf []byte) (Executable, error) {

	arch, err := elf_arch(elf)
	if err != nil {
		return nil, err
	}

	endianness, err := elf_endianness(elf)
	if err != nil {
		return nil, err
	}

	section_header_offset := elf_section_header_offset(elf, arch, endianness)

	return Elf{
		arch:       arch,
		endianness: endianness,
	}, nil
}

func elf_arch(elf []byte) (uint, error) {
	arch_index := 4
	if elf[arch_index] == 1 {
		return 32, nil
	} else if elf[arch_index] == 2 {
		return 64, nil
	} else {
		return 0, errors.New("invalid ELF file")
	}
}

func elf_endianness(elf []byte) (string, error) {
	endianness_index := 5
	if elf[endianness_index] == 1 {
		return "little", nil
	} else if elf[endianness_index] == 2 {
		return "big", nil
	} else {
		return "", errors.New("invalid ELF file")
	}
}

// try to generalize this so you don't have to keep repeating yourself
func elf_section_header_offset(elf []byte, arch uint, endianness string) uint {
	var section_header_offset []byte
	var sect_hdr_offst_loc int
	var sect_hrd_offst_size int
	if arch == 64 {
		sect_hdr_offst_loc = 0x28
		sect_hrd_offst_size = 8
	} else if arch == 32 {
		sect_hdr_offst_loc = 0x20
		sect_hrd_offst_size = 4
	}
	section_header_offset = elf[sect_hdr_offst_loc : sect_hdr_offst_loc+sect_hrd_offst_size]
	if endianness == "big" {
		if arch == 64 {
			return uint(binary.BigEndian.Uint64(section_header_offset))
		}
		if arch == 32 {
			return uint(binary.BigEndian.Uint32(section_header_offset))
		}
	}
	if endianness == "little" {
		if arch == 64 {
			return uint(binary.LittleEndian.Uint64(section_header_offset))
		}
		if arch == 32 {
			return uint(binary.LittleEndian.Uint32(section_header_offset))
		}
	}
	return 0 // this will never be reached
}
