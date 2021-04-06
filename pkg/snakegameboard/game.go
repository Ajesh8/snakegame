package snakegameboard

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

func reset() {
	term.Sync() // cosmestic purpose
}

func StartGame(height int, width int) (int, string) {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	game := new(gameState)
	game.IntiliazeGameState(height, width)
keyPressListenerLoop:
	for {
		fmt.Printf("Score:%d Snake Length:%d\n", game.score, game.length)
		fmt.Println("User arrow keys to move. Press ESC to quit")
		game.printBoard()
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
			case term.KeyArrowUp:
				reset()
				game.currentRound = "up"
			case term.KeyArrowDown:
				reset()
				game.currentRound = "down"
			case term.KeyArrowLeft:
				reset()
				game.currentRound = "left"
			case term.KeyArrowRight:
				reset()
				game.currentRound = "right"
			default:
				reset()
				fmt.Println("Use arrow keys to move the snake")
				continue
			}
		case term.EventError:
			panic(ev.Err)
		}
		game.handleSnakeMovement()
		if game.gameOver != 0 {
			break keyPressListenerLoop
		}
	}
	var message string
	if game.gameOver == 1 {
		message = "Your snake hit a wall and died of concussion."
	} else if game.gameOver == 2 {
		message = "Your snake died by eating itself like a maniac."
	} else {
		message = "You quit the game."
	}
	return game.score, message
}
