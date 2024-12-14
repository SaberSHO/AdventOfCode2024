package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

func main() {
	encoding.Register()
	scn, err := tcell.NewScreen()
	scn.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	x := 10
	y := 10
	for {
		scn.Show()
		ev := scn.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'n' {
				y++
				x++
				scn.Clear()
				scn.SetContent(y, x, rune('#'), []rune(""), tcell.StyleDefault)
			}
			if ev.Rune() == 'x' {
				scn.Fini()
				os.Exit(0)
			}
		}
	}
}
