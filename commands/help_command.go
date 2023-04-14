package commands

import "fmt"

func help(cfg *ReplConfig, _ string) error {
	fmt.Println("Available commands:")
	for _, command := range GetCLICommands() {
		fmt.Printf("\t%s: %s\n", command.Name, command.Description)
	}
	return nil
}
