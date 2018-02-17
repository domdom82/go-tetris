package main

import tl "github.com/JoelOtter/termloop"
import "log"
import "bufio"
import (
	"os"
	"math/rand"
	"time"
)

const fps = 10

// Termloop stuff
var game *tl.Game
var score *Score
var playerlevel int

func gameOverScreen() {
	e := NewGameOver()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorWhite,
	})
	game.Screen().SetLevel(level)
	level.AddEntity(e)
}

func startScreen() {
	e := NewTitle()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorWhite,
	})
	game.Screen().SetLevel(level)
	level.AddEntity(e)
}


func gameScreen() {
	level := NewLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorBlack,
	})
	game.Screen().SetLevel(level)

	// Create playing field
	gameArea := NewGameArea()

	// Create first tile
	tile := NewTile(gameArea)

	// Create score
	score = NewScore()

	gameArea.AddToLevel(level)
	tile.AddToLevel(level)
	game.Screen().AddEntity(score)
}

func main() {
	// Set up RNG
	rand.Seed(time.Now().UTC().UnixNano())

	// Set up logging
	logfile := "go-tetris.log"
	file, err := os.Create(logfile)
	if err != nil {
		log.Fatal("Could not open log file ", logfile)
	}
	writer := bufio.NewWriter(file)
	log.SetOutput(writer)
	defer file.Close()
	defer writer.Flush()

	// Game setup
	playerlevel = 0
	game = tl.NewGame()
	game.Screen().SetFps(fps)
	startScreen()
	game.Start()
}
