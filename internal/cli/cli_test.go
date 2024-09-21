package cli

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/0x4D5352/pokedexcli/internal/config"
)

func TestCommandMap(t *testing.T) {
	commands := getCommands()
	cases := []string{
		"help",
		"exit",
		"map",
		"mapb",
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			command, ok := commands[c]
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if command.name != c {
				t.Errorf("expected command name %s, got command name %s", c, command.name)
				return
			}
		})
	}
}

func TestExecuteCommand(t *testing.T) {
	cfg := &config.Config{}
	for name := range getCommands() {
		if name == "exit" {
			// gotta handle anything with os.exit seperately!
			continue
		}
		err := ExecuteCommand(name, cfg)
		if err != nil {
			t.Errorf("expected %s, got %v", name, err)
			return
		}
	}
}

func TestExit(t *testing.T) {
	if os.Getenv("EXIT_CALL") == "1" {
		cfg := &config.Config{}
		commandExit(cfg)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestExit")
	cmd.Env = append(os.Environ(), "EXIT_CALL=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 0", err)
}

func TestMapBack(t *testing.T) {
	cfg := &config.Config{}
	err := commandMapBack(cfg)
	if err == nil {
		t.Errorf("expected error, got previous location %v", cfg.Previous)
		return
	}
	if err.Error() != "already at beginning of list!" {
		t.Errorf("got unexpected error %v", err)
		return
	}
	return
}
