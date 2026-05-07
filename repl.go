package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		reader.Scan()

		if err := reader.Err(); err != nil {
			fmt.Println(err)
			return
		}

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands[commandName]

		if exists {
			if err := command.callback(); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
