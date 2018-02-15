package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Border struct {
	*tl.Rectangle
	xoffset int
	yoffset int
}


func (b *Border) Draw(s *tl.Screen) {
	screenWidth,screenHeight := s.Size()
	myWidth,myHeight := b.Size()

	b.SetPosition(screenWidth / 2 - myWidth / 2 + b.xoffset, screenHeight / 2 - myHeight / 2 + b.yoffset)



	b.Rectangle.Draw(s)
}

//func (snake *Border) Tick(event tl.Event) {
//	if event.Type == tl.EventKey { // Is it a keyboard event?
//		switch event.Key { // If so, switch on the pressed key.
//		case tl.KeyArrowRight:
//		snake.xoffset--
//		case tl.KeyArrowLeft:
//		snake.xoffset--
//		case tl.KeyArrowUp:
//		snake.yoffset--
//		case tl.KeyArrowDown:
//		snake.yoffset++
//		case tl.KeySpace:
//			//TESTING
//			//food.Reset()
//		}
//	}
//}

func NewBorder(xoffset int, yoffset int, width int, height int, color tl.Attr) *Border {
	t := new(Border)

	t.xoffset = xoffset
	t.yoffset = yoffset
	t.Rectangle = tl.NewRectangle(1,1,width,height,color)

	return t
}
