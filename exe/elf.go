package exe

import (
	"iago/isa"
	// "iago/suffixtree"
	"encoding/binary"
	"errors"
)

type Elf struct {
	arch       uint // either 32 or 64
	endianness string
	isa        isa.ISA
	contents   []byte
	// section_header_offset uint
}

type elfHeaderEntry struct {
	offset32 uint
	offset64 uint
	size32 uint
	size64 uint
}

// map from all necessary elf header field names to their locations and sizes
// fields have been renamed for clarity
var elfHeader = map[string]elfHeaderEntry{
	"arch": {0x04, 0x04, 1, 1},
	"endianness": {0x05, 0x05, 1, 1},
	"isa": {0x12, 0x12, 2, 2},
	"entry_point": {0x18, 0x18, 4, 8},
	"section_header_table_offset": {0x20, 0x28, 4, 8},
	"section_header_table_entry_size": {0x2e, 0x3a, 2, 2},
	"section_header_table_num_entries": {0x30, 0x3c, 2, 2},
	"section_header_table_names_index": {0x32, 0x3e, 2, 2},
}

func (e Elf) foo() {

}

func NewElf(elfContents []byte) (Executable, error) {

	arch, err := elfArch(elfContents)
	if err != nil {
		return nil, err
	}

	endianness, err := elfEndianness(elfContents)
	if err != nil {
		return nil, err
	}

	elf := Elf{
		arch: arch,
		endianness: endianness,
		contents: elfContents,
		isa: nil,
	}

	elf.setISA()

	return elf, nil
}



func elfArch(elfContents []byte) (uint, error) {
	archOffset := 4
	if elfContents[archOffset] == 1 {
		return 32, nil
	} else if elfContents[archOffset] == 2 {
		return 64, nil
	} else {
		return 0, errors.New("invalid ELF file")
	}
}

func elfEndianness(elfContents []byte) (string, error) {
	endiannessOffset := 5
	if elfContents[endiannessOffset] == 1 {
		return "little", nil
	} else if elfContents[endiannessOffset] == 2 {
		return "big", nil
	} else {
		return "", errors.New("invalid ELF file")
	}
}

func (e *Elf) headerValue(field string) uint {
	var offset uint
	var size uint
	fieldInfo := elfHeader[field]
	if e.arch == 32 {
		offset = fieldInfo.offset32
		size = fieldInfo.size32
	} else if e.arch == 64 {
		offset = fieldInfo.offset64
		size = fieldInfo.size64
	}

	if e.endianness == "big" {
		return uint(binary.BigEndian.Uint64(e.contents[offset:offset + size]))
	} else if e.endianness == "little" {
		return uint(binary.LittleEndian.Uint64(e.contents[offset:offset + size]))
	}
	return 0 // this will never be reached
}

func (e *Elf) setISA() (isa.ISA, error) {

	// maps the value present in the elf file to an ISA
	var supportedISAs = map[uint]isa.ISA{
		0x03: isa.X86{},
		0x3e: isa.X86{},
	}

	isa, ok := supportedISAs[e.headerValue("isa")]
	if ok {
		return isa, nil
	}

	return nil, errors.New("unsupported instruction set")
}

