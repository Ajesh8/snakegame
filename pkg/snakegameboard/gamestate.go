package snakegameboard

import (
	"fmt"
	"math/rand"
	"time"
)

type coordinate struct {
	x int
	y int
}

type gameState struct {
	board          [][]int
	currentRound   string
	score          int
	length         int
	snakeHead      coordinate
	snakeTail      coordinate
	foodCoordinate coordinate
	moveGrid       [][]string
}

func (p *gameState) GetRandomCoordinate(height int, width int) coordinate {
	rand.Seed(time.Now().UnixNano())
	var result coordinate
	result.x = rand.Intn(width)
	result.y = rand.Intn(height)
	return result
}

func (p *gameState) setFoodCoordinate(food coordinate) {
	for p.board[food.y][food.x] == 1 {
		food = p.GetRandomCoordinate(len(p.board), len(p.board[0]))
	}
	p.foodCoordinate = food
}

func (p *gameState) IntiliazeGameState(height int, width int) {
	p.board = make([][]int, height)
	p.moveGrid = make([][]string, height)
	for i := range p.board {
		p.board[i] = make([]int, width)
		p.moveGrid[i] = make([]string, width)
	}
	p.score = 0
	p.length = 1
	p.snakeHead = coordinate{
		x: width / 2,
		y: height / 2,
	}
	p.snakeTail = coordinate{
		x: width / 2,
		y: height / 2,
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			p.board[i][j] = 0
		}
	}
	p.board[p.snakeHead.y][p.snakeHead.x] = 1
	p.setFoodCoordinate(p.GetRandomCoordinate(height, width))
	p.board[p.foodCoordinate.y][p.foodCoordinate.x] = 2
}

func (p *gameState) printBoard() {
	for i := 0; i < len(p.board); i++ {
		for j := 0; j < len(p.board[i]); j++ {
			if p.board[i][j] == 0 {
				fmt.Print(".")
			}
			if p.board[i][j] == 1 {
				fmt.Print("X")
			}
			if p.board[i][j] == 2 {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func nextCoordinate(x int, y int, dir string) (int, int) {
	if dir == "up" {
		return x, y - 1
	} else if dir == "down" {
		return x, y + 1
	} else if dir == "left" {
		return x - 1, y
	}
	return x + 1, y
}
func (p *gameState) handleSnakeMovement() {
	p.moveGrid[p.snakeHead.y][p.snakeHead.x] = p.currentRound
	p.snakeHead.x, p.snakeHead.y = nextCoordinate(p.snakeHead.x, p.snakeHead.y, p.currentRound)
	p.board[p.snakeHead.y][p.snakeHead.x] = 1
	if p.foodCoordinate.x != p.snakeHead.x || p.foodCoordinate.y != p.snakeHead.y {
		p.board[p.snakeTail.y][p.snakeTail.x] = 0
		p.snakeTail.x, p.snakeTail.y = nextCoordinate(p.snakeTail.x, p.snakeTail.y, p.moveGrid[p.snakeTail.y][p.snakeTail.x])
		p.board[p.snakeTail.y][p.snakeTail.x] = 1
	} else {
		p.length++
		p.setFoodCoordinate(p.GetRandomCoordinate(len(p.board), len(p.board[0])))
		p.board[p.foodCoordinate.y][p.foodCoordinate.x] = 2
		p.score++
	}
}
