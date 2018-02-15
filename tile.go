package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"golang.org/x/tools/refactor/rename"
)


const baseDelay = 20

type Type int

const (
	LONG Type = iota
	SQUARE
	T
	SQUIGGLY_LEFT
	SQUIGGLY_RIGHT
)

// A tile consists of 4 parts, each of which can be a block or nil
// When all parts are nil, the tile is considered cleared and can be removed from the game
// It is also itself an Entity just for wiring up the keyboard controls and base position
type Tile struct {
	*tl.Entity
	part1 *Block
	part2 *Block
	part3 *Block
	part4 *Block
	delay int
	tileType Type
	rotation int
}


// Tick for a tile
func (tile *Tile) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		x,y := tile.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			tile.SetPosition(x+2,y)
		case tl.KeyArrowLeft:
			tile.SetPosition(x-2,y)
		case tl.KeyArrowUp:
			tile.rotate()
		case tl.KeyArrowDown:
			tile.SetPosition(x,y+1)
		case tl.KeySpace:
			//TESTING
			//food.Reset()
		}
	}
}

// Draw for a tile
func (tile *Tile) Draw(screen *tl.Screen) {
	x,y := tile.Position()

	tile.delay--

	if tile.delay <= 0 {
		tile.delay = baseDelay - playerlevel
		y++
	}

	tile.SetPosition(x,y)
	tile.part1.SetPosition(x+tile.part1.xoffset, y+tile.part1.yoffset)
	tile.part2.SetPosition(x+tile.part2.xoffset, y+tile.part2.yoffset)
	tile.part3.SetPosition(x+tile.part3.xoffset, y+tile.part3.yoffset)
	tile.part4.SetPosition(x+tile.part4.xoffset, y+tile.part4.yoffset)

}

func (tile *Tile) rotate() {
	tile.rotation += 90 % 360

	//tbd adjust part offsets based on rotation and type
}

func (tile *Tile) Collide(collision tl.Physical) {
	switch collision.(type) {
	//case *Food:
	//	f := collision.(*Food)
	//	score.updateScore(f.score * len(snake.body))
	//	snake.grow = 5
		//f.Reset(snake)
	}
}

func (tile *Tile) AddToLevel(level tl.Level) {
	level.AddEntity(tile.part1)
	level.AddEntity(tile.part2)
	level.AddEntity(tile.part3)
	level.AddEntity(tile.part4)
}

func NewTile() *Tile {
	t := new(Tile)
	t.Entity = tl.NewEntity(1, 1, 1, 1)



	//Choose random tile
	choice := rand.Intn(4)
	switch choice {
	case 0:
		// long piece
		t.tileType = LONG
	case 1:
		// square
		t.tileType = SQUARE
	case 2:
		// t-piece
		t.tileType = T
	case 3:
		// squiggly left
		t.tileType = SQUIGGLY_LEFT
	case 4:
		// squiggly right
		t.tileType = SQUIGGLY_RIGHT
	}

	t.delay = baseDelay - playerlevel
	t.rotation = 0
	return t
}
