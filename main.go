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
		fmt.Printf("User Input: %s\n", input)
	}
}
