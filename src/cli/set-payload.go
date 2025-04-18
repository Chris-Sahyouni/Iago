package cli

import (
	"encoding/binary"
	"errors"
	"iago/src/global"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type SetPayload struct{ args Args }

func (s SetPayload) ValidArgs() bool {
	// if len(s.args) == 0 {
	// 	return true
	// 	// will open interactive editor in this case
	// }

	if len(s.args) != 1 {
		return false
	}

	_, ok := s.args["default"]
	return ok

}

func (s SetPayload) Execute(globalState *global.GlobalState) error {
	// if len(s.args) == 0 {
	// 	fmt.Println("Would open interactive editor in this case")
	// 	return nil
	// }

	fileName := s.args["default"]

	paddingLength, err := paddingSizeFromFileName(fileName)
	if err != nil {
		return err
	}

	globalState.CurrentPayload.PaddingLength = paddingLength

	contents, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	if len(contents) < paddingLength {
		return errors.New("file is smaller than specified padding")
	}

	globalState.CurrentPayload.PaddingLength = paddingLength

	arch := globalState.CurrentFile.Arch()
	endianness := globalState.CurrentFile.Endianness()
	chain := ReadChainFromFileContents(contents, arch, endianness, paddingLength)

	globalState.CurrentPayload.Chain = chain

	return nil
}

func paddingSizeFromFileName(fileName string) (int, error) {
	padSubstrIdx := strings.Index(fileName, "pad")

	if padSubstrIdx == -1 {
		return 0, nil
	}

	bytesOfPaddingStr := ""
	runeArrFileName := []rune(fileName)
	for i := padSubstrIdx + 3; i < len(runeArrFileName); i++ {
		r := runeArrFileName[i]
		if unicode.IsDigit(r) {
			bytesOfPaddingStr += string(r)
		}
	}
	if bytesOfPaddingStr == "" {
		return -1, errors.New("\"pad\" substring found in filename but padding size not specified")
	}
	bytesOfPadding, err := strconv.Atoi(bytesOfPaddingStr)
	if err != nil {
		return -1, err
	}
	return bytesOfPadding, nil
}

func ReadChainFromFileContents(contents []byte, arch uint, endianness string, paddingLength int) []uint {
	bytesPerAddr := arch / 8
	var byteorder binary.ByteOrder
	if endianness == "little" {
		byteorder = binary.LittleEndian
	} else {
		byteorder = binary.BigEndian
	}

	chainBytes := contents[paddingLength:]

	chain := make([]uint, 0)
	for i := 0; i < len(chainBytes); i += int(bytesPerAddr) {
		addrBytes := chainBytes[i : i+int(bytesPerAddr)]
		if arch == 32 {
			addr := byteorder.Uint32(addrBytes)
			chain = append(chain, uint(addr))
		} else {
			addr := byteorder.Uint64(addrBytes)
			chain = append(chain, uint(addr))
		}
	}
	return chain
}
