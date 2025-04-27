package term

import (
	"fmt"
	"os"
	"golang.org/x/term"
)

type Terminal struct{
	term *term.Terminal
}

func RawTerminal() Terminal {
	stdInFd := int(os.Stdin.Fd())
	_, err := term.MakeRaw(stdInFd)
	if err != nil {
		panic(err)
	}
	// defer term.Restore(stdInFd, oldState)
	return Terminal{
		term.NewTerminal(os.Stdin, "> "),
	}
}

func (t *Terminal) ReadLine() (string, error) {
	return t.term.ReadLine()
}

// necessary when in raw terminal mode
func Println(a ...any) {
	fmt.Print(a...)
	fmt.Print("\r\n")
}
