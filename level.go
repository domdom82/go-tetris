package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Level struct {
	*tl.BaseLevel
}

func NewLevel(bg tl.Cell) *Level {
	level := new(Level)
	level.BaseLevel = tl.NewBaseLevel(bg)
	return level
}

func (l *Level) Tick(ev tl.Event) {

	if ev.Type == tl.EventNone {
		for _, e := range l.Entities {
			e.Tick(ev)
		}
	}
		l.BaseLevel.Tick(ev)

}

