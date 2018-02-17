package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Block struct {
	*tl.Rectangle
	tile *Tile
	xoffset int
	yoffset int
	prevX int
	prevY int
}


func (block *Block) Collide(collision tl.Physical) {

	switch collision.(type) {
	case *Border:
		x,y := block.Position()
		cx,cy := collision.Position()
		block.tile.resetPosition()

		// we've hit the bottom border
		if y == cy && x != cx {
			block.tile.movable = false
			// launch new tile
			newTile := NewTile(block.tile.gameArea)
			newTile.AddToLevel(game.Screen().Level())
		}
	//case *Block:
	//	x,y := block.Position()
	//	_,cy := collision.Position()
	//	block.tile.resetPosition()
	//
	//	if y == cy && x == block.prevX {
	//		block.tile.movable = false
	//		// launch new tile
	//		newTile := NewTile(block.tile.gameArea)
	//		newTile.AddToLevel(game.Screen().Level())
	//	}
	}


}

func (block *Block) SetPosition(x int, y int) {
	block.prevX, block.prevY = block.Position()
	block.Rectangle.SetPosition(x,y)
}

func (block *Block) Draw(screen *tl.Screen) {
	if block.tile.invisible == false {
		block.Rectangle.Draw(screen)
	}
}


func NewBlock(tile *Tile, xoffset int, yoffset int) *Block {
	b := new(Block)
	b.Rectangle = tl.NewRectangle(1,1, 2, 1, tl.ColorRed)
	b.xoffset = xoffset
	b.yoffset = yoffset
	b.tile = tile

	return b
}
