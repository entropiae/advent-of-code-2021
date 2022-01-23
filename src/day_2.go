package main

import (
	"bufio"
	"fmt"
	"os"
)

type command struct {
	direction string
	value     int
}

func readCommands(path string) ([]command, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var commands []command
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value := scanner.Text()
		commands = append(commands, parseCommand(value))
	}
	return commands, nil
}

func parseCommand(raw string) command {
	var direction string
	var value int
	fmt.Sscanf(raw, "%s %d", &direction, &value)

	return command{direction: direction, value: value}
}

func computeFirstStep(commands []command) {
	depth := 0
	position := 0

	for _, c := range commands {
		switch c.direction {
		case "forward":
			position += c.value
		case "down":
			depth += c.value
		case "up":
			depth -= c.value
		default:
			fmt.Println("Invalid command")
		}
	}
	fmt.Printf("Depth: %d, position: %d, aggregate: %d\n", depth, position, depth*position)
}

func computeSecondStep(commands []command) {
	depth := 0
	position := 0
	aim := 0

	for _, c := range commands {
		switch c.direction {
		case "down":
			aim += c.value
		case "up":
			aim -= c.value
		case "forward":
			position += c.value
			depth += aim * c.value
		default:
			fmt.Println("Invalid command")
		}
	}
	fmt.Printf("Depth: %d, position: %d, aggregate: %d\n", depth, position, depth*position)
}

func main() {
	file := os.Args[1]
	commands, err := readCommands(file)

	if err != nil {
		os.Exit(-1)
	}
	computeFirstStep(commands)
	computeSecondStep(commands)
}
