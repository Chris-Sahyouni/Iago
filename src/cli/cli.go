package cli

import (
	"errors"
	"iago/src/global"
	"strings"
)

type Command interface {
	Execute(*global.GlobalState) error
	ValidArgs() bool
}

type Args = map[string]string

/*
For this project any flags (anything beginning with - or --) take an argument directly following it.
Any non-flag argument is considered the commands singular default argument (i.e <path> in load <path>)
and will be represented by "default": "<value>" in the args dictionary.
*/
func ParseLine(line string) (Command, error) {

	inputs := strings.Split(strings.TrimSpace(line), " ")
	cmdName := inputs[0]
	flagsAndArgs := inputs[1:]
	args := Args{}

	if len(inputs) == 0 {
		return nil, errors.New("")
	}

	if len(inputs) == 1 {
		return commandDispatch(cmdName, args)
	}

	flagState := false
	defaultSet := false
	var flag string
	for _, val := range flagsAndArgs {
		if val[0] == '-' { // flag

			if flagState {
				return nil, errors.New("invalid arguments")
			}

			flag = val
			flagState = true
		} else { // default or flag value
			if flagState { // flag value
				args[flag] = val
				flagState = false
			} else { // default
				if defaultSet {
					return nil, errors.New("invalid arguments")
				}
				args["default"] = val
				defaultSet = true
			}
		}
	}
	return commandDispatch(cmdName, args)
}

func commandDispatch(cmdName string, args Args) (Command, error) {

	var cmd Command

	switch cmdName {
	case "help":
		cmd = Help{args}
	case "quit":
		cmd = Quit{args}
	case "load":
		cmd = Load{args}
	case "stat":
		cmd = Stat{args}
	case "cat-target":
		cmd = CatTarget{args}
	case "set-target":
		cmd = SetTarget{args}
	case "rop":
		cmd = Rop{args}
	default:
		return nil, errors.New("unrecognized command")
	}

	return cmd, nil
}