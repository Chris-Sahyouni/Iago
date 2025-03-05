package exe

import (
	"encoding/binary"
	"errors"
	"iago/isa"
	"fmt"
)

type Elf struct {
	arch       uint // either 32 or 64
	endianness string // either "big" or "little"
	isa        isa.ISA
	contents   []byte
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
	"endianness": {0x05, 0x05, 1,  1},
	"isa": {0x12, 0x12, 1, 1}, // technically this field is 2 bytes, but the 2nd byte is only used for two obscure ISAs
	"entry point": {0x18, 0x18, 4, 8},
	"program header table offset": {0x1c, 0x20, 4, 8},
	"program header table entry size": {0x2a, 0x36, 2, 2},
	"program header table num entries": {0x2c, 0x38, 2, 2},

	// not sure we need the section header yet
	// "section_header_table_offset": {0x20, 0x28, 4, 8},
	// "section_header_table_entry_size": {0x2e, 0x3a, 2, 2},
	// "section_header_table_num_entries": {0x30, 0x3c, 2, 2},
	// "section_header_table_names_index": {0x32, 0x3e, 2, 2},
}

func (e *Elf) Info()  {
	fmt.Println("  File Type: ELF")
	fmt.Println("  Arch:", e.arch)
	fmt.Println("  ISA:", e.isa.Name())
	fmt.Println("  Endianness:", e.endianness)
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

	elf := &Elf{
		arch: arch,
		endianness: endianness,
		contents: elfContents,
		isa: nil,
	}

	err = elf.setISA()
	if err != nil {
		return nil, err
	}

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

	value := e.contents[offset:offset + size]

	if size == 1 {
		return uint(value[0])
	}

	var byteOrder binary.ByteOrder

	if e.endianness == "big" {
		byteOrder = binary.BigEndian
	} else if e.endianness == "little" {
		byteOrder = binary.LittleEndian
	}

	if size == 2 {
		return uint(byteOrder.Uint16(value))
	} else if size == 4 {
		return uint(byteOrder.Uint32(value))
	} else if size == 8 {
		return uint(byteOrder.Uint64(value))
	}

	return 0 // this will never be reached
}

func (e *Elf) setISA() error {

	// maps the value present in the elf file to an ISA
	var supportedISAs = map[uint]isa.ISA{
		0x03: isa.X86{},
		0x3e: isa.X86{},
	}

	isa, ok := supportedISAs[e.headerValue("isa")]
	if ok {
		e.isa = isa
		return nil
	}

	return errors.New("unsupported instruction set")
}

