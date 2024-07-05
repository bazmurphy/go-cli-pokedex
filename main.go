package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	startRepl()
}

func startRepl() {
	config := &Config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
	}

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		inputText := reader.Text()

		command, exists := getCommands()[inputText] // (!) this is how to lookup on a map that is returned from a function
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of (the next) 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
	}
}

func commandHelp(config *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(config *Config) error {
	os.Exit(0)
	return nil
}

func commandMap(config *Config) error {
	err := getLocationArea(config.next, config)
	if err != nil {
		return err
	}

	return nil
}

func commandMapB(config *Config) error {
	if config.previous == "" {
		return errors.New("cannot go any further back")
	}

	err := getLocationArea(config.previous, config)
	if err != nil {
		return err
	}

	return nil
}

// TODO (!) i don't like having to pass the config all the way down here...
// but we need to update the config object with the new values that come back in the response
func getLocationArea(url string, config *Config) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var locationAreaData LocationAreaResponse
	err = json.Unmarshal(body, &locationAreaData)
	if err != nil {
		return err
	}

	// fmt.Println("DEBUG | count:", locationAreaData.Count)
	// fmt.Println("DEBUG | next:", locationAreaData.Next)
	// fmt.Println("DEBUG | previous:", locationAreaData.Previous)

	for _, area := range locationAreaData.Results {
		fmt.Println(area.Name)
	}

	// TODO (!) per the above comments
	config.next = locationAreaData.Next
	config.previous = locationAreaData.Previous

	return nil
}
