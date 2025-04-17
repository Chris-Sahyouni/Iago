package exe

import (
	"encoding/hex"
	"fmt"
	"iago/src/isa"
	"os"
	"testing"
	"reflect"
)

var testBinaries = map[string][]byte{}

// reads in test binaries
func setup() {
	testPath := "../../test_programs/bin/"
	dirEntries, err := os.ReadDir(testPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(dirEntries) == 0 {
		fmt.Println("Compile the test programs before running tests")
		os.Exit(1)
	}
	for _, f := range dirEntries {
		name := f.Name()
		fmt.Println(name)
		fileContents, err := os.ReadFile(testPath + name)
		if err != nil {
			fmt.Println("Error reading test binary")
			os.Exit(1)
		}
		testBinaries[name] = fileContents
	}
	// added in error case
	corruptFile, err := hex.DecodeString("df118e9f9cc24bd9a1c989d57fc28976568d1f49b2e19352f81d48ef1b41f2595e1b6ec3e7553deb680d76b3aeb7e7faa576baec526553be4bc5c1c5900c2450851151ddef4031d69a30843750753215ec22a811fe02e73fee80df7db60d0bff6d80f43ac3d2116fc230f59e8d463f66f7442bc85f2a717f92b4c6ab6db347e3")
	if err != nil {
		fmt.Println("Error decoding hex string")
		os.Exit(1)
	}
	testBinaries["corrupt"] = corruptFile
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func TestNewElf(t *testing.T) {
	var expectedResults = map[string]struct {
		Arch uint
		End  string
		Isa  isa.ISA
	}{
		"square32": {Arch: 32, End: "little", Isa: isa.X86{}},
		"square64": {Arch: 64, End: "little", Isa: isa.X86{}},
	}
	for name, contents := range testBinaries {
		expected := expectedResults[name]
		actual, err := NewElf(contents)
		if name == "corrupt" {
			if err == nil {
				t.Error("Should have failed on corrupt file")
			}
		} else {
			if actual.(*Elf).arch != expected.Arch {
				t.Fail()
			}
			if actual.(*Elf).endianness != expected.End {
				t.Fail()
			}
			if actual.(*Elf).isa != expected.Isa {
				t.Fail()
			}
		}
	}
}

// this test might fail if you recompile the test binaries
func TestFieldValue(t *testing.T) {
	var expectedResults = map[string]struct {
		EntryPnt     uint
		PHdrEntrySz  uint
		PHdrVirtAddr uint // of the first entry of the Program Header Table
		Flags        uint // of the first entry of the Program Header Table
	}{
		"square32": {EntryPnt: 0x1070, PHdrEntrySz: 32, PHdrVirtAddr: 0x34, Flags: 0x4},
		"square64": {EntryPnt: 0x1040, PHdrEntrySz: 56, PHdrVirtAddr: 0x40, Flags: 0x4},
	}
	for name, contents := range testBinaries {
		expected := expectedResults[name]
		elf, err := NewElf(contents)
		if err != nil {
			t.Error(err)
		}
		if name == "corrupt" {
			continue
		}
		value, err := elf.(*Elf).fieldValue("entry point", elfHeader, 0)
		if err != nil {
			t.Error(err)
		}
		if value != expected.EntryPnt {
			t.Fail()
		}
		value, err = elf.(*Elf).fieldValue("program header table entry size", elfHeader, 0)
		if err != nil {
			t.Error(err)
		}
		if value != expected.PHdrEntrySz {
			t.Fail()
		}
		PHhdrOffset, err := elf.(*Elf).fieldValue("program header table offset", elfHeader, 0)
		if err != nil {
			t.Error(err)
		}
		value, err = elf.(*Elf).fieldValue("virtual address", programHeaderEntry, PHhdrOffset)
		if err != nil {
			t.Error(err)
		}
		if value != expected.PHdrVirtAddr {
			t.Fail()
		}
		value, err = elf.(*Elf).fieldValue("flags", programHeaderEntry, PHhdrOffset)
		if err != nil {
			t.Error(err)
		}
		if value != expected.Flags {
			t.Fail()
		}
	}
}

func TestLocateExecutableSegments(t *testing.T) {

	var expectedResults = map[string][]segment{
		"square32": {
			segment{
				Offset: 0x1000,
				VAddr:  0x1000,
				Size:   0x294,
			},
		},
		"square64": {
			segment{
				Offset: 0x1000,
				VAddr:  0x1000,
				Size:   0x1e5,
			},
		},
	}

	for name, contents := range testBinaries {
		expected := expectedResults[name]
		elf, err := NewElf(contents)
		if err != nil {
			t.Error(err)
		}
		actual, err := elf.(*Elf).locateExecutableSegments()
		if err != nil {
			t.Error(err)
		}

		t.Log(actual)

		for i := range len(actual) {
			if actual[i] != expected[i] {
				t.Fail()
			}
		}
	}
}

func TestInstructionStream(t *testing.T) {

	testContents := []byte("abcdefz")

	testElf := Elf{
		arch: 0,
		endianness: "little",
		isa: isa.TestISA{},
		contents: testContents,
		programHeaderTableOffset: 0,
		reverseInstructionTrie: nil,
	}

	testSegments := []segment{
		{ // a,b,c
			VAddr: 0,
			Offset: 0,
			Size: 3,
		},
		{ // d
			VAddr: 20,
			Offset: 3,
			Size: 1,
		},
		{ // empty
			VAddr: 30,
			Offset: 5,
			Size: 0,
		},
	}

	expected := []isa.Instruction{
		{
			Vaddr: 0,
			Op: "a",
		},
		{
			Vaddr: 1,
			Op: "b",
		},
		{
			Vaddr: 2,
			Op: "c",
		},
		{
			Vaddr: 3,
			Op: "d",
		},
	}

	actual := testElf.InstructionStream(testSegments)

	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}

}