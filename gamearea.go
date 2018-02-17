package main

import (
	tl "github.com/JoelOtter/termloop"
)

type GameArea struct {
	*tl.Entity
	left *Border
	right *Border
	bottom *Border
}

func (a *GameArea) Draw(s *tl.Screen) {
	screenWidth,screenHeight := s.Size()
	areaWidth,areaHeight := a.Size()

	a.SetPosition(screenWidth / 2 - areaWidth / 2, screenHeight / 2 - areaHeight /2)

	x,y := a.Position()

	a.bottom.SetPosition(x+a.bottom.xoffset,y+a.bottom.yoffset)
	a.left.SetPosition(x+a.left.xoffset,y+a.left.yoffset)
	a.right.SetPosition(x+a.right.xoffset,y+a.right.yoffset)

}

func (a *GameArea) AddToLevel(level tl.Level) {
	level.AddEntity(a)
	level.AddEntity(a.bottom)
	level.AddEntity(a.left)
	level.AddEntity(a.right)
}


func NewGameArea() *GameArea {
	a := new(GameArea)

	a.Entity = tl.NewEntity(1, 1, 24, 20)

	a.bottom = NewBorder(0, 20, 24, 1, tl.ColorBlue)
	a.left   = NewBorder(0, 0, 2, 20, tl.ColorBlue)
	a.right  = NewBorder(22, 0, 2, 20, tl.ColorBlue)


	return a
}
