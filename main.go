package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	running := true
	inputScanner := bufio.NewScanner(os.Stdin)
	for running {
		fmt.Print("pokedex > ")
		inputScanner.Scan()
		input := inputScanner.Text()
		fmt.Println()
		err := executeCommand(input)
		if err != nil {
			fmt.Println(err)
		}
	}
}
