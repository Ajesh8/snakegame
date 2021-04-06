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
	gameOver       int
}

func (p *gameState) GetUnoccupiedRandomCoordinate() coordinate {
	rand.Seed(time.Now().UnixNano())
	var emptyCell []coordinate
	for i := 0; i < len(p.board); i++ {
		for j := 0; j < len(p.board[0]); j++ {
			if p.board[i][j] == 0 {
				emptyCell = append(emptyCell, coordinate{
					x: j,
					y: i,
				})
			}
		}
	}
	return emptyCell[rand.Intn(len(emptyCell))]
}

func (p *gameState) IntiliazeGameState(height int, width int) {
	p.gameOver = 0
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
	p.board[p.snakeHead.y][p.snakeHead.x] = 3
	p.foodCoordinate = p.GetUnoccupiedRandomCoordinate()
	p.board[p.foodCoordinate.y][p.foodCoordinate.x] = 2
}

func (p *gameState) printBoard() {
	for i := -1; i <= len(p.board); i++ {
		for j := -1; j <= len(p.board[0]); j++ {
			if i == -1 || i == len(p.board) {
				fmt.Print(" -")
				continue
			}
			if j == -1 || j == len(p.board[0]) {
				fmt.Print(" |")
				if i == -1 {
					fmt.Print("| ")
				}
				continue
			}
			if p.board[i][j] == 0 {
				fmt.Print("  ")
			}
			if p.board[i][j] == 1 {
				fmt.Print("o ")
			}
			if p.board[i][j] == 2 {
				fmt.Print("X ")
			}
			if p.board[i][j] == 3 {
				fmt.Print("O ")
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

func checkBoundaryCollision(x int, y int, width int, height int) bool {
	if x < 0 || y < 0 || x == width || y == height {
		return true
	}
	return false
}
func (p *gameState) handleSnakeMovement() {
	p.moveGrid[p.snakeHead.y][p.snakeHead.x] = p.currentRound
	p.board[p.snakeHead.y][p.snakeHead.x] = 1
	p.snakeHead.x, p.snakeHead.y = nextCoordinate(p.snakeHead.x, p.snakeHead.y, p.currentRound)
	if checkBoundaryCollision(p.snakeHead.x, p.snakeHead.y, len(p.board[0]), len(p.board)) {
		p.gameOver = 1
		return
	}
	if p.board[p.snakeHead.y][p.snakeHead.x] == 1 {
		p.gameOver = 2
		return
	}
	p.board[p.snakeHead.y][p.snakeHead.x] = 3
	if p.foodCoordinate.x != p.snakeHead.x || p.foodCoordinate.y != p.snakeHead.y {
		p.board[p.snakeTail.y][p.snakeTail.x] = 0
		p.snakeTail.x, p.snakeTail.y = nextCoordinate(p.snakeTail.x, p.snakeTail.y, p.moveGrid[p.snakeTail.y][p.snakeTail.x])
		if p.snakeTail.x != p.snakeHead.x || p.snakeTail.y != p.snakeHead.y {
			p.board[p.snakeTail.y][p.snakeTail.x] = 1
		}
	} else {
		p.length++
		p.foodCoordinate = p.GetUnoccupiedRandomCoordinate()
		p.board[p.foodCoordinate.y][p.foodCoordinate.x] = 2
		p.score++
	}
}
