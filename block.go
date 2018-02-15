package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)

type Block struct {
	*tl.Entity
	xoffset int
	yoffset int
}


func (block *Block) Collide(collision tl.Physical) {
	switch collision.(type) {
	//case *Food:
	//	f := collision.(*Food)
	//	score.updateScore(f.score * len(snake.body))
	//	snake.grow = 5
		//f.Reset(snake)
	}
}


func NewBlock(xoffset int, yoffset int) *Block {
	b := new(Block)
	b.Entity = tl.NewEntity(1, 1, 2, 1)
	b.Entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '█'})
	b.xoffset = xoffset
	b.yoffset = yoffset

	return b
}
