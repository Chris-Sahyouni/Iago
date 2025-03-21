package cli

import (
	"bytes"
	"errors"
	"iago/exe"
	"iago/global"
	"os"
)

type Load struct{ args Args }

func (l Load) ValidArgs() bool {
	if len(l.args) != 1 {
		return false
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
	exeConstructors := map[string]func([]byte) (exe.Executable, error){
		"elf": exe.NewElf,
	}
	fileType, err := determineFileType(fileBytes)
	if err != nil {
		return err
	}
	constructor := exeConstructors[fileType]
	newExecutable, err := constructor(fileBytes)
	if err != nil {
		return err
	}
	globalState.CurrentFile = newExecutable

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
