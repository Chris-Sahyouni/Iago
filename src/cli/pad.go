package cli

import (
	"errors"
	"iago/src/global"
	"iago/src/term"
	"os"
	"strconv"
	"strings"
)

type Pad struct{ args Args }

func (p Pad) ValidArgs() bool {

	if len(p.args) != 1 {
		return false
	}

	_, ok := p.args["default"]
	return ok
}

func (p Pad) Execute(globalState *global.GlobalState) error {
	paddingLength, err := strconv.Atoi(p.args["default"])
	if err != nil {
		return errors.New("failed to convert inputted padding length to integer")
	}

	padding := strings.Repeat("=", paddingLength)
	globalState.CurrentPayload.PaddingLength = paddingLength

	outName := "rop_chain_pad" + strconv.Itoa(paddingLength)
	outFile, err := os.OpenFile(outName, os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer outFile.Close()

	outFile.Write([]byte(padding))
	WriteChainToFile(globalState.CurrentPayload.Chain, globalState.CurrentFile.Arch(), globalState.CurrentFile.Endianness(), outFile)

	return nil
}

func (Pad) Help() {
	term.Println("    pad <bytes>" + strings.Repeat(" ", SPACE_BETWEEN-len("pad <bytes>")) + "Generate a new payload by prepending <bytes> number of bytes of padding to the current payload.")
}
