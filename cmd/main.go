package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Ajesh8/snakegame/pkg/snakegameboard"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Print("Height and Width for board not provided. Please try again.")
	} else {
		height, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("error during string to int conversion:%v", err)
		}
		width, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("error during string to int conversion:%v", err)
		}
		if height > 40 {
			fmt.Println("Box height might exceed the commandline. Setting height to 40")
			height = 40
		} else if height < 10 {
			height = 10
			fmt.Println("Box height is too small for the snake game. Setting it to 10")
		}
		if width > 60 {
			width = 60
			fmt.Println("Box width might exceed the commandline. Setting height to 60")
		} else if width < 10 {
			width = 10
			fmt.Println("Box width is too small for the snake game. Setting it to 10")
		}
		score, message := snakegameboard.StartGame(height, width)
		fmt.Printf("Game Over. %s Your final score is:%d\n", message, score)
	}
}
