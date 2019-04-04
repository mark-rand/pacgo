package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func loadMaze() error {
	mazePath := "maze01.txt"

	f, err := os.Open(mazePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return nil
}

var maze []string

func printScreen() {
	for _, line := range maze {
		fmt.Println(line)
	}
}

func main() {
	// initialize game

	// load maze
	err := loadMaze()
	if err != nil {
		log.Printf("Error loading maze: %v", err)
	}

	// game loop
	for {
		// update screen
		printScreen()

		// get input

		// process movement

		// check game over

		// Temp: break infinite loop
		break

		// repeat
	}
}
