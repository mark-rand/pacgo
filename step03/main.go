package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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

func readInput() (string, error) {
	buffer := make([]byte, 10)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	}

	return "", nil
}

func initialize() {
	cbreakTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbreakTerm.Stdin = os.Stdin

	err := cbreakTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cooked mode terminal: %v\n", err)
	}
}

func main() {
	// initialize game
	initialize()
	defer cleanup()

	// load maze
	err := loadMaze()
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
	}

	// game loop
	for {
		// update screen
		printScreen()

		// get input
		input, err := readInput()
		if err != nil {
			log.Printf("Error reading input: %v", err)
			break
		}

		// process movement

		fmt.Println(input)

		// check game over
		if input == "ESC" {
			break
		}

		// repeat
	}
}
