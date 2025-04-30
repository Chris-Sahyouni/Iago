package cli

import (
	"bytes"
	"errors"
	"iago/src/exe"
	"iago/src/global"
	"iago/src/term"
	"os"
	"strings"
)

type Load struct{ args Args }

func (l Load) ValidArgs() bool {
	if len(l.args) != 1 && len(l.args) != 2 {
		return false
	}

	if len(l.args) == 2 {
		_, ok := l.args["--thumb"]
		if !ok {
			return false
		}
	}

	_, ok := l.args["default"]
	return ok
}

func (l Load) Execute(globalState *global.GlobalState) error {
	path := l.args["default"]

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	exeConstructors := map[string]func([]byte, Args) (exe.Executable, error){
		"elf": exe.NewElf,
	}
	fileType, err := determineFileType(fileBytes)
	if err != nil {
		return err
	}
	constructor := exeConstructors[fileType]
	newExecutable, err := constructor(fileBytes, l.args)
	if err != nil {
		return err
	}

	globalState.CurrentFile = newExecutable

	// invalidate current payload on loading new file
	globalState.CurrentPayload = struct{PaddingLength int; Chain []uint}{
		0, nil,
	}


	newExecutable.Info()

	return nil
}

func determineFileType(fileBytes []byte) (string, error) {
	elfMagic := []byte{'\x7f', '\x45', '\x4c', '\x46'}
	if bytes.Equal(fileBytes[:4], elfMagic) {
		return "elf", nil
	}

	return "", errors.New("unrecognized file format")
}


func (Load) Help() {
	term.Println("    load <path> [OPTIONS]" + strings.Repeat(" ", SPACE_BETWEEN-len("load <path> [OPTIONS]")) + "Sets the current file for analysis")
	term.Println("        --thumb" + strings.Repeat(" ", SPACE_BETWEEN-len("    --thumb")) + "Targets thumb mode for ARM binaries")
}