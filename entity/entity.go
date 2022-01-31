package entity

import (
	"fmt"
)

type Entity struct {
	X      int
	Y      int
	Symbol rune
}

func (e *Entity) MoveEntity(r rune) {
	switch r {
	case 'w':
		e.Y--
	case 'a':
		e.X--
	case 's':
		e.Y++
	case 'd':
		e.X++
	default:
		fmt.Println("Unknown key")
	}
}

func (e *Entity) DisplayEntityPosition() {
	fmt.Printf("%c at %d, %d\n", e.Symbol, e.X, e.Y)
}
