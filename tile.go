package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)


const baseDelay = 20

type Type int

const (
	LONG Type = iota
	SQUARE
	T
	Z_LEFT
	Z_RIGHT
	L_LEFT
	L_RIGHT
)

// A tile consists of 4 parts, each of which can be a block or nil
// When all parts are nil, the tile is considered cleared and can be removed from the game
// It is also itself an Entity just for wiring up the keyboard controls and base position
type Tile struct {
	*tl.Entity
	gameArea *GameArea
	part1 *Block
	part2 *Block
	part3 *Block
	part4 *Block
	delay int
	delayCounter int
	invisibleCounter int
	tileType Type
	movable bool
	invisible bool
	rotation int
	x int
	y int
	prevX int
	prevY int
}

func (tile *Tile) move(x int, y int) {
	if tile.movable {
		tile.SetPosition(x,y)
		xabs,yabs := tile.Entity.Position()
		tile.part1.SetPosition(xabs+tile.part1.xoffset, yabs+tile.part1.yoffset)
		tile.part2.SetPosition(xabs+tile.part2.xoffset, yabs+tile.part2.yoffset)
		tile.part3.SetPosition(xabs+tile.part3.xoffset, yabs+tile.part3.yoffset)
		tile.part4.SetPosition(xabs+tile.part4.xoffset, yabs+tile.part4.yoffset)
	}
}

func (tile *Tile) resetDelay() {
	tile.delayCounter = tile.delay
}

func (tile *Tile) resetPosition() {
	tile.move(tile.prevX, tile.prevY)
}

func (tile *Tile) SetPosition(x int, y int) {
	// Set position relative to game area
	xArea,yArea := tile.gameArea.Position()
	tile.x = x
	tile.y = y
	tile.Entity.SetPosition(x + xArea, y + yArea)
}

func (tile *Tile) Position() (int,int){
	// Return position relative to game area
	return tile.x,tile.y
}

// Tick for a tile
func (tile *Tile) Tick(event tl.Event) {
	x,y := tile.Position()
	tile.prevX = x
	tile.prevY = y

	if tile.invisible == true {
		tile.invisibleCounter--
		if tile.invisibleCounter <= 0 {
			tile.invisible = false
		}
	}

	if tile.movable == true {
		tile.delayCounter--
		if tile.delayCounter <= 0 {
			tile.resetDelay()
			y++
		}
	}

	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			x+=2
		case tl.KeyArrowLeft:
			x-=2
		case tl.KeyArrowUp:
			tile.rotate()
		case tl.KeyArrowDown:
			y++
			tile.resetDelay()
		case tl.KeySpace:
			tile.delay = 1 //drop tile
			tile.resetDelay()
		}
	}

	tile.move(x,y)
}

func (tile *Tile) rotate() {
	tile.rotation += 90 % 360

	//tbd adjust part offsets based on rotation and type
}

func (tile *Tile) AddToLevel(level tl.Level) {
	level.AddEntity(tile)
	level.AddEntity(tile.part1)
	level.AddEntity(tile.part2)
	level.AddEntity(tile.part3)
	level.AddEntity(tile.part4)
}

func NewTile(area *GameArea) *Tile {
	t := new(Tile)
	t.gameArea = area
	t.Entity = tl.NewEntity(1, 1, 1, 1)
	areaWidth,_ := area.Size()
	var xoff,yoff int
	//Choose random tile
	choice := rand.Intn(7)
	//DEBUG
	choice = 0
	switch choice {
	case 0:
		t.tileType = LONG
		t.part1 = NewBlock(t, 0, 0)
		t.part2 = NewBlock(t, 0, 1)
		t.part3 = NewBlock(t, 0, 2)
		t.part4 = NewBlock(t, 0, 3)
		xoff = areaWidth / 2
		yoff = -2
	case 1:
		t.tileType = SQUARE
	case 2:
		t.tileType = T
	case 3:
		t.tileType = Z_LEFT
	case 4:
		t.tileType = Z_RIGHT
	case 5:
		t.tileType = L_LEFT
	case 6:
		t.tileType = L_RIGHT
	}

	t.delay = baseDelay - playerlevel
	t.delayCounter = t.delay
	t.invisibleCounter = 1
	t.rotation = 0
	t.movable = true
	t.invisible = true
	t.move(xoff,yoff)
	return t
}
