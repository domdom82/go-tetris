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
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorBlack,
	})
	game.Screen().SetLevel(level)

	// Create playfield
	area := NewGameArea()


	// Create score
	score = NewScore()

	level.AddEntity(area)
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
	game = tl.NewGame()
	game.Screen().SetFps(fps)
	startScreen()
	game.Start()
}
