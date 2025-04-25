package cli

import (
	"encoding/binary"
	"errors"
	"iago/src/global"
	"os"
)

type Rop struct{ args Args }

func (r Rop) ValidArgs() bool {

	// might do something like this later to generalize arg management for all commands
	// allowed := []struct{
	// 	short string
	// 	long string
	// }{
	// 	{short: "-f", long: "--file"},
	// 	{short: "-t", long: "--target"},
	// 	{short: "-o", long: "--out"},
	// }

	// for _, flag := range r.args {
	// 	if !slices.ContainsFunc(allowed, func(e struct{short string; long string}) bool {return flag == e.short || flag == e.long}) {
	// 		return false
	// 	}
	// }

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

	gadgetAddrs, err := globalState.CurrentFile.Rop(globalState.CurrentTarget.Contents)
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
		var gAddrBytes []byte
		if arch == 32 {
			byteorder.PutUint32(gAddrBytes, uint32(gAddr))
			outFile.Write(gAddrBytes)
		}
		if arch == 64 {
			byteorder.PutUint64(gAddrBytes, uint64(gAddr))
			outFile.Write(gAddrBytes)
		}
	}
}