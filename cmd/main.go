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
		width, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("error during string to int conversion:%v", err)
		}
		snakegameboard.StartGame(height, width)
	}
}
