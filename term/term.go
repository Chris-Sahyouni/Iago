package term

import (
	"fmt"
	"io"
	"os"
	"golang.org/x/term"
)

type Terminal struct{
	term *term.Terminal
	oldState *term.State
}

func RawTerminal() Terminal {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	return Terminal{
		term.NewTerminal(os.Stdin, "> "),
		oldState,
	}
}

func (t *Terminal) ReadLine() (string, error) {
	line, err := t.term.ReadLine()
	if err == io.EOF { // handles ctrl+c, ctrl+d
		return "quit", nil
	}

	return line, nil
}

func (t *Terminal) Restore() {
	term.Restore(int(os.Stdin.Fd()), t.oldState)
}

// necessary when in raw terminal mode
func Println(a ...any) {
	fmt.Print(a...)
	fmt.Print("\r\n")
}
