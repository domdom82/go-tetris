package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Field uint8

const (
	EMPTY Field = iota
	BLOCK
	BORDER
)

type GameArea struct {
	*tl.Entity
	field [][]Field
}


func (t *GameArea) Tick(event tl.Event) {
}

func (a *GameArea) Draw(s *tl.Screen) {
	screenWidth,screenHeight := s.Size()
	myWidth,myHeight := a.Size()

	a.SetPosition(screenWidth / 2 - myWidth / 2, screenHeight / 2 - myHeight / 2)

	offsetX,offsetY := a.Position()
	col := tl.ColorBlack
	for y := range a.field {
		for x := range a.field[y] {
			switch a.field[y][x] {
			case EMPTY:
				col = tl.ColorBlack
			case BLOCK:
				col = tl.ColorWhite
			case BORDER:
				col = tl.ColorBlue
			}

			s.RenderCell(x + offsetX, y + offsetY, &tl.Cell{
				Bg: col,
			})
		}
	}

}


func NewGameArea() *GameArea {
	a := new(GameArea)

	// playing field is 10x20 blocks with borders left/right and top/bottom adding 2 blocks each
	width := 22
	height := 22

	a.Entity = tl.NewEntity(1, 1, width, height)
	a.field = make([][]Field, height)
	for y  := range a.field {
		a.field[y] = make([]Field, width)
	}

	// set borders
	for y := range a.field {
		for x := range a.field[y] {
			if x == 0 || x == width -1 || y == 0 || y == height -1 {
				a.field[y][x] = BORDER
			} else {
				a.field[y][x] = EMPTY
			}

		}
	}

	return a
}
