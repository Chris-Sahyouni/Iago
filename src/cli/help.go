package cli

import (
	"github.com/Chris-Sahyouni/iago/src/global"
	"github.com/Chris-Sahyouni/iago/src/term"
	"strings"
)

const SPACE_BETWEEN = 24

type Help struct{ args Args }

func (h Help) ValidArgs() bool {
	return len(h.args) == 0
}

func (h Help) Execute(_ *global.GlobalState) error {
	term.Println("Commands:")
	Help{}.Help()
	Load{}.Help()
	Stat{}.Help()
	CatTarget{}.Help()
	SetTarget{}.Help()
	CatPayload{}.Help()
	SetPayload{}.Help()
	Find{}.Help()
	Rop{}.Help()
	Pad{}.Help()
	Quit{}.Help()

	return nil
}

func (Help) Help() {
	term.Println("    help" + strings.Repeat(" ", SPACE_BETWEEN-len("help")) + "View this help menu")
}