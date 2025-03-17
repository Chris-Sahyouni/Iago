package exe

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"iago/isa"
	"fmt"
)

type Elf struct {
	arch       uint // either 32 or 64
	endianness string // either "big" or "little"
	isa        isa.ISA
	contents   []byte
	programHeaderTableOffset uint
	ExecutableSegments []Segment
}

type elfField struct {
	offset32 uint
	offset64 uint
	size32 uint
	size64 uint
}

type Segment struct {
	VAddr uint
	Offset uint
	Size uint
}


var elfHeader = map[string]elfField {
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

var programHeaderEntry = map[string]elfField {
	"segment type": {0x00, 0x00, 4, 4},
	"flags": {0x18, 0x04, 4, 4},
	"segment offset": {0x04, 0x08, 4, 8},
	"virtual address": {0x08, 0x10, 4, 8},
	"file size": {0x10, 0x20, 4, 8},
	"mem size": {0x14, 0x28, 4, 8},
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
		programHeaderTableOffset: 0,
		isa: nil,
	}



	err = elf.setISA()
	if err != nil {
		return nil, err
	}

	err = elf.locateExecutableSegments()
	if err != nil {
		return nil, err
	}

	return elf, nil
}



func elfArch(elfContents []byte) (uint, error) {
	archOffset := 4

	if len(elfContents) < archOffset {
		return 0, errors.New("invalid ELF file")
	}

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

	if len(elfContents) < endiannessOffset {
		return "", errors.New("invalid ELF file")
	}

	if elfContents[endiannessOffset] == 1 {
		return "little", nil
	} else if elfContents[endiannessOffset] == 2 {
		return "big", nil
	} else {
		return "", errors.New("invalid ELF file")
	}
}

func (e *Elf) fieldValue(field string, targetHeader map[string]elfField, baseOffset uint) (uint, error) {
	var offset uint
	var size uint

	fieldInfo := targetHeader[field]
	if e.arch == 32 {
		offset = fieldInfo.offset32
		size = fieldInfo.size32
	} else if e.arch == 64 {
		offset = fieldInfo.offset64
		size = fieldInfo.size64
	}

	offset += baseOffset

	if len(e.contents) < int(offset + size) {
		return 0, errors.New("value offset outside file bounds")
	}

	value := e.contents[offset:offset + size]

	if size == 1 {
		return uint(value[0]), nil
	}

	var byteOrder binary.ByteOrder

	if e.endianness == "big" {
		byteOrder = binary.BigEndian
	} else if e.endianness == "little" {
		byteOrder = binary.LittleEndian
	}

	if size == 2 {
		return uint(byteOrder.Uint16(value)), nil
	} else if size == 4 {
		return uint(byteOrder.Uint32(value)), nil
	} else if size == 8 {
		return uint(byteOrder.Uint64(value)), nil
	}

	return 0, nil // this will never be reached
}

func (e *Elf) setISA() error {

	// maps the value present in the elf file to an ISA
	var supportedISAs = map[uint]isa.ISA{
		0x03: isa.X86{},
		0x3e: isa.X86{},
	}

	value, err := e.fieldValue("isa", elfHeader, 0)
	if err != nil {
		return err
	}

	isa, ok := supportedISAs[value]
	if ok {
		e.isa = isa
		return nil
	}

	return errors.New("unsupported instruction set")
}

func (e *Elf) locateExecutableSegments() error {
	var segments []Segment

	programHeaderTableOffset, err := e.fieldValue("program header table offset", elfHeader, 0)
	if err != nil {
		return err
	}
	programHeaderTableEntrySize, err := e.fieldValue("program header table entry size", elfHeader, 0)
	if err != nil {
		return err
	}
	numEntries, err := e.fieldValue("program header table num entries", elfHeader, 0)
	if err != nil {
		return err
	}
	for i := range numEntries {
		entryOffset := programHeaderTableOffset + (i * programHeaderTableEntrySize)
		flags, err := e.fieldValue("flags", programHeaderEntry, entryOffset)
		if err != nil {
			return err
		}
		var executableFlag uint = 0x1
		if flags == executableFlag {
			segmentOffset, err := e.fieldValue("segment offset", programHeaderEntry, entryOffset)
			if err != nil {
				return err
			}
			virtualAddress, err := e.fieldValue("virtual address", programHeaderEntry, entryOffset)
			if err != nil {
				return err
			}
			sizeInFile, err := e.fieldValue("file size", programHeaderEntry, entryOffset)
			if err != nil {
				return err
			}

			segments = append(segments, Segment{
				VAddr: virtualAddress,
				Offset: segmentOffset,
				Size: sizeInFile,
			})

		}
	}
	e.ExecutableSegments = segments
	return nil
}


func (e *Elf) InstructionStream() []isa.Instruction {
	var instructionStream []isa.Instruction
	instructionSize := e.isa.InstructionSize()
	for _, segment := range e.ExecutableSegments {
		segmentContents := e.contents[segment.Offset:segment.Offset + segment.Size]
		for i := 0; i < len(segmentContents); i += instructionSize {
			newInstruction := isa.Instruction{
				// make sure this is correct for big endian programs too
				Op: hex.EncodeToString(segmentContents[i:i+instructionSize]),
				Vaddr: segment.VAddr + uint(i),
			}
			instructionStream = append(instructionStream, newInstruction)
		}
	}
	return instructionStream
}