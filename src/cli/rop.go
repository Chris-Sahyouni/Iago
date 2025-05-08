package cli

import (
	"encoding/binary"
	"errors"
	"github.com/Chris-Sahyouni/iago/src/global"
	"github.com/Chris-Sahyouni/iago/src/term"
	"os"
	"strings"
)

type Rop struct{ args Args }

func (r Rop) ValidArgs() bool {

	if len(r.args) == 0 {
		return true
	}

	if len(r.args) == 1 {
		_, ok := r.args["-o"]
		return ok
	}

	return false
}

func (r Rop) Execute(globalState *global.GlobalState) error {
	if globalState.CurrentFile == nil {
		return errors.New("no file loaded. Run load <path>")
	}
	if globalState.CurrentTarget == struct {
		Title    string
		Contents string
	}{"", ""} {
		return errors.New("no target payload specified. Run set-target <path>")
	}

	outName, ok := r.args["-o"]
	if !ok {
		outName = "rop_chain"
	}

	outFile, err := os.OpenFile(outName, os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer outFile.Close()

	currFile := globalState.CurrentFile
	gadgetAddrs, err := currFile.ReverseInstructionTrie().Rop(globalState.CurrentTarget.Contents, currFile.Isa())
	if err != nil {
		return err
	}

	endianness := globalState.CurrentFile.Endianness()
	arch := globalState.CurrentFile.Arch()

	WriteChainToFile(gadgetAddrs, arch, endianness, outFile)

	return nil
}

func WriteChainToFile(chain []uint, arch uint, endianness string, outFile *os.File) {
	var byteorder binary.ByteOrder
	if endianness == "big" {
		byteorder = binary.BigEndian
	} else {
		byteorder = binary.LittleEndian
	}

	for _, gAddr := range chain {
		if arch == 32 {
			gAddrBytes := make([]byte, 4)
			byteorder.PutUint32(gAddrBytes, uint32(gAddr))
			outFile.Write(gAddrBytes)
		}
		if arch == 64 {
			gAddrBytes := make([]byte, 8)
			byteorder.PutUint64(gAddrBytes, uint64(gAddr))
			outFile.Write(gAddrBytes)
		}
	}
}

func (Rop) Help() {
	term.Println("    rop [OPTIONS]" + strings.Repeat(" ", SPACE_BETWEEN-len("rop [OPTIONS]")) + "Generate a ROP chain for the target file and payload")
	term.Println("        -o FILE" + strings.Repeat(" ", SPACE_BETWEEN - len("    -o FILE")) + "Specify output file (optional)")

}