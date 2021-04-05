package snakegameboard

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

func reset() {
	term.Sync() // cosmestic purpose
}

func StartGame(height int, width int) {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	game := new(gameState)
	game.IntiliazeGameState(height, width)
	fmt.Println("Enter any key to see their ASCII code or press ESC button to quit")
keyPressListenerLoop:
	for {
		fmt.Printf("Score:%d Snake Length:%d\n", game.score, game.length)
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
	}
}
