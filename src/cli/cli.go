package cli

import (
	"errors"
	"iago/src/global"
	"strings"
)

type Command interface {
	Execute(*global.GlobalState) error
	ValidArgs() bool
	Help()
}

type Args = map[string]string


/*
For this project any flags beginning with - take an argument directly following it.
Any flags beginning with -- do not take an argument
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

	expectingArg := false
	defaultSet := false
	var flag string
	for _, val := range flagsAndArgs {
		if val[0] == '-' { // flag

			if expectingArg {
				return nil, errors.New("invalid arguments")
			}

			if val[1] == '-' { // non-arg-taking flag
				args[val] = ""
			} else { // arg-taking flag
				flag = val
				expectingArg = true
			}

		} else { // default or flag arg
			if expectingArg { // flag arg
				args[flag] = val
				expectingArg = false
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
	case "cat-payload":
		cmd = CatPayload{args}
	case "set-payload":
		cmd = SetPayload{args}
	case "find":
		cmd = Find{args}
	case "rop":
		cmd = Rop{args}
	case "pad":
		cmd = Pad{args}
	default:
		return nil, errors.New("unrecognized command")
	}

	return cmd, nil
}