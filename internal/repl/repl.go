package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/0x4D5352/pokedexcli/internal/cli"
	"github.com/0x4D5352/pokedexcli/internal/config"
	"github.com/0x4D5352/pokedexcli/internal/pokecache"
)

func StartRepl() {
	running := true
	inputScanner := bufio.NewScanner(os.Stdin)
	cfg := config.Config{
		Next:     "",
		Previous: "",
		Cache:    *pokecache.NewCache(time.Minute * 5),
	}
	for running {
		fmt.Print("Pokedex > ")
		inputScanner.Scan()
		input := cleanInput(inputScanner.Text())
		commandName := input[0]
		err := cli.ExecuteCommand(commandName, &cfg)
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
