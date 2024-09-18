package repl

import (
	"bufio"
	"fmt"
	"github.com/0x4D5352/pokedexcli/internal/cli"
	"github.com/0x4D5352/pokedexcli/internal/config"
	"os"
	"strings"
)

func StartRepl() {
	running := true
	inputScanner := bufio.NewScanner(os.Stdin)
	config := config.Config{
		Next:     "",
		Previous: "",
	}
	for running {
		fmt.Print("Pokedex > ")
		inputScanner.Scan()
		input := cleanInput(inputScanner.Text())
		commandName := input[0]
		fmt.Println()
		err := cli.ExecuteCommand(commandName, config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
