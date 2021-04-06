package snakegameboard

import (
	"fmt"
	"math/rand"
	"time"
)

type coordinate struct { // struct version of a cell position.
	x int
	y int
}

// gameState contains variables that define the current state of the game.
type gameState struct {
	board          [][]int    // board stores which cell is currently empty, occupied by snake body or food.
	currentRound   string     // Stores the current direction the snake is heading
	score          int        // Stores the score of user
	length         int        // Length of snake
	snakeHead      coordinate // position of snake's head
	snakeTail      coordinate // position of snake's tail
	foodCoordinate coordinate // position of food particle
	moveGrid       [][]string // stores the direction snake was heading when it was on that cell.
	gameOver       int
}

// getUnoccupiedRandomCoordinate randomly picks an empty cell for putting a food unit.
func (p *gameState) GetUnoccupiedRandomCoordinate() coordinate {
	rand.Seed(time.Now().UnixNano())
	var emptyCell []coordinate
	for i := 0; i < len(p.board); i++ {
		for j := 0; j < len(p.board[0]); j++ {
			if p.board[i][j] == 0 { // if cell is empty, add to slice
				emptyCell = append(emptyCell, coordinate{
					x: j,
					y: i,
				})
			}
		}
	}
	return emptyCell[rand.Intn(len(emptyCell))] // pick a random index from slice.
}

// This function initiliazes the gameState variable when game starts.
func (p *gameState) IntiliazeGameState(height int, width int) {
	p.gameOver = 0
	p.board = make([][]int, height)
	p.moveGrid = make([][]string, height)
	for i := range p.board {
		p.board[i] = make([]int, width)
		p.moveGrid[i] = make([]string, width) // Initialising board and moveGrid.
	}
	p.score = 0
	p.length = 1
	p.snakeHead = coordinate{ // Putting snake in the center of board.
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
	p.board[p.snakeHead.y][p.snakeHead.x] = 3 // cell occupied by snakehead has value 3. Other snake parts have value 1.
	p.foodCoordinate = p.GetUnoccupiedRandomCoordinate()
	p.board[p.foodCoordinate.y][p.foodCoordinate.x] = 2 // cell occupied by food unit has value 2.
}

// printing the current board on console.
func (p *gameState) printBoard() {
	for i := -1; i <= len(p.board); i++ {
		for j := -1; j <= len(p.board[0]); j++ {
			if i == -1 || i == len(p.board) { // upper and lower border.
				fmt.Print(" -")
				continue
			}
			if j == -1 || j == len(p.board[0]) { // left and right border.
				fmt.Print(" |")
				if i == -1 {
					fmt.Print("| ")
				}
				continue
			}
			if p.board[i][j] == 0 { // empty cells
				fmt.Print("  ")
			}
			if p.board[i][j] == 1 { // snake body parts
				fmt.Print("o ")
			}
			if p.board[i][j] == 2 { // food particle
				fmt.Print("X ")
			}
			if p.board[i][j] == 3 { // snake head
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

// check if snake is colliding with the wall or not
func checkBoundaryCollision(x int, y int, width int, height int) bool {
	if x < 0 || y < 0 || x == width || y == height {
		return true
	}
	return false
}

// handling snake movement across board.
func (p *gameState) handleSnakeMovement() {
	p.moveGrid[p.snakeHead.y][p.snakeHead.x] = p.currentRound // setting current direction.
	p.board[p.snakeHead.y][p.snakeHead.x] = 1
	p.snakeHead.x, p.snakeHead.y = nextCoordinate(p.snakeHead.x, p.snakeHead.y, p.currentRound) // Getting next cell for snake head based on current direction.
	if checkBoundaryCollision(p.snakeHead.x, p.snakeHead.y, len(p.board[0]), len(p.board)) {    // if snake collided with the border wall.
		p.gameOver = 1
		return
	}
	if p.board[p.snakeHead.y][p.snakeHead.x] == 1 { // if snake head collides with other body part.
		p.gameOver = 2
		return
	}
	p.board[p.snakeHead.y][p.snakeHead.x] = 3
	if p.foodCoordinate.x != p.snakeHead.x || p.foodCoordinate.y != p.snakeHead.y { // if snake head is not eating food unit.
		p.board[p.snakeTail.y][p.snakeTail.x] = 0
		p.snakeTail.x, p.snakeTail.y = nextCoordinate(p.snakeTail.x, p.snakeTail.y, p.moveGrid[p.snakeTail.y][p.snakeTail.x])
		if p.snakeTail.x != p.snakeHead.x || p.snakeTail.y != p.snakeHead.y {
			p.board[p.snakeTail.y][p.snakeTail.x] = 1
		}
	} else { // if snake eat the food unit.
		p.length++
		p.foodCoordinate = p.GetUnoccupiedRandomCoordinate()
		p.board[p.foodCoordinate.y][p.foodCoordinate.x] = 2
		p.score++
	}
}
