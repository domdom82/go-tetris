package main

import tl "github.com/JoelOtter/termloop"
import "log"
import "bufio"
import (
	"os"
	"math/rand"
	"time"
	"golang.org/x/tools/go/gcimporter15/testdata"
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
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorBlack,
	})
	game.Screen().SetLevel(level)

	// Create playfield
	bottom := NewBorder(0, 10, 21, 1, tl.ColorBlue)
	left := NewBorder(-10, 0, 1, 22, tl.ColorBlue)
	right := NewBorder(+10, 0, 1, 22, tl.ColorBlue)


	// Create score
	score = NewScore()

	level.AddEntity(bottom)
	level.AddEntity(left)
	level.AddEntity(right)
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
