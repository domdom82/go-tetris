package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Border struct {
	*tl.Rectangle
	xoffset int
	yoffset int
}

func NewBorder(xoffset int, yoffset int, width int, height int, color tl.Attr) *Border {
	t := new(Border)

	t.xoffset = xoffset
	t.yoffset = yoffset
	t.Rectangle = tl.NewRectangle(1,1,width,height,color)

	return t
}
