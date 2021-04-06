package snakegameboard

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

func reset() {
	err := term.Sync()
	if err != nil {
		fmt.Print(fmt.Errorf("could not reset termbox-go:%w", err))
	}
}

// StartGame function controls the flow of the game.
func StartGame(height int, width int) (int, string) {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	game := new(gameState) // Creating an instance of gameState to maintain the state of game.
	game.IntiliazeGameState(height, width)
keyPressListenerLoop: // Loop will run unless either ESC key is pressed or the game is over by either hitting the bounday or hitting body of snake.
	for {
		fmt.Printf("Score:%d Snake Length:%d\n", game.score, game.length)
		fmt.Println("User arrow keys to move. Press ESC to quit")
		game.printBoard() // Print score, instructions and the complete board after every turn.
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
			case term.KeyArrowUp:
				reset()
				game.currentRound = "up" // currentRound sets the current direction the snake is headed.
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
		game.handleSnakeMovement() // Handling snake movement.
		if game.gameOver != 0 {
			break keyPressListenerLoop //Game over condition.
		}
	}
	var message string
	if game.gameOver == 1 { // Ending messages for different scenarios in game over event.
		message = "Your snake hit a wall and died of concussion."
	} else if game.gameOver == 2 {
		message = "Your snake died by eating itself like a maniac."
	} else {
		message = "You quit the game."
	}
	return game.score, message
}
