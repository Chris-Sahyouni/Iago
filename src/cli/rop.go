package cli

import (
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
	if globalState.TargetPayload == struct{Title string; Contents string}{"", ""} {
		return errors.New("no target payload specified. Run set-target <path>")
	}

	outputDir, ok := r.args["-o"]
	if !ok {
		outputDir = "iago_generated_payloads"
	}

	err := os.Mkdir(outputDir, 0755)
	if err != nil {
		return err
	}

	return nil
}
