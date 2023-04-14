package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/geekayush1812/pokedex/pokeapi"
)

type Cli struct {
	Name        string
	Description string
	Callback    func(*ReplConfig, string) error
}

type ReplConfig struct {
	HttpClient pokeapi.Client
}

func (cfg *ReplConfig) Start() {
	fmt.Println("Welcome to Pokedex CLI. Type 'help' to see all available commands.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		parts := strings.Split(input, " ")

		var command, args string = input, ""

		switch len(parts) {
		case 1:
			command = parts[0]

		case 2:
			command = parts[0]
			args = parts[1]

		default:
			fmt.Println("Invalid command")
			continue
		}

		cliCommands := GetCLICommands()

		if inputCommand, ok := cliCommands[command]; ok {
			if error := inputCommand.Callback(cfg, args); error != nil {
				fmt.Println(error)
			}

		} else {
			fmt.Println("Invalid command")
		}
	}
}

func GetCLICommands() map[string]Cli {
	return map[string]Cli{
		"help": {
			Name:        "help",
			Description: "Show all available commands",
			Callback:    help,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex CLI",
			Callback:    exit,
		},
		"map": {
			Name:        "map",
			Description: "Show the 20 location area of the pokemon world, subsequent calls will show the next 20 locations",
			Callback:    mapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Show the previous 20 location area of the pokemon world, subsequent calls will show the previous 20 locations",
			Callback:    mapbCommand,
		},
		"explore": {
			Name:        "explore",
			Description: "Show the pokemon names that can be found in the given location area",
			Callback:    exploreCommand,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch the pokemon with the given name",
			Callback:    catchCommand,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Show the stats of the pokemon with the given name",
			Callback:    inspectCommand,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Show the list of all the caught pokemon",
			Callback:    pokedexCommand,
		},
	}
}
