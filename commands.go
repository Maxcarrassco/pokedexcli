package main

import (
	"math/rand"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Maxcarrassco/pokedexcli/internal"
	"github.com/Maxcarrassco/pokedexcli/models"
)

var Mapping = map[string]models.Command{
	"exit": models.Command{
		Description: "Exits the Pokedex",
		CallbackFn: commandExit,
	},
	"map": models.Command{
		Description: "Displays the names of 20 location areas in the Pokemon world",
		CallbackFn: commandMap,
	},
	"mapb": models.Command{
		Description: "Displays the names of 20 location areas backward in the Pokemon world",
		CallbackFn: commandMapb,
	},
	"explore": models.Command{
		Description: "Lists all the Pokémon in a given area",
		CallbackFn: commandExplore,
	},
	"catch": models.Command{
		Description: "Catch a Pokemon and add it to the user's Pokedex",
		CallbackFn: commandCatch,
	},
	"inspect": models.Command{
		Description: "Print the name, height, weight, stats and type(s) of the Pokemon",
		CallbackFn: commandInspect,
	},
	"pokedex": models.Command{
		Description: "Print the name of Pokemon the user have caught",
		CallbackFn: commandPokedex,
	},
}

var UserPokemon = make(map[string]models.PokedexCatch)


var pokedexLocation = models.PokedexLocation{}
var pokedexLocationArea = models.PokedexLocationArea{}
var pokedexCatch = models.PokedexCatch{}

func commandExit(args string) error {
	fmt.Println("Thanks for playing Pokedex")
	os.Exit(0)
	return nil
}

func CommandHelp() {
	fmt.Println("Welcome to Pokedex\n\nUsage:\n\n")
	for k, v := range Mapping {
		fmt.Println(k+":", v.Description + "\n")
	}
	fmt.Println("help:", "Displays the help page of the Pokedex")
}


func commandMap(args string) error {
	if pokedexLocation.Next == nil {
	   url := fmt.Sprintf("%s/%s/?offset=0&limit=20", internal.BASE_URL, "location")
	   err := internal.GetRequest(url, &pokedexLocation)
	   if err != nil {
		return err
	   }
	} else {
	  err := internal.GetRequest(*pokedexLocation.Next, &pokedexLocation)
	  if err != nil {
		return err
	  }
        }
	for _, v := range pokedexLocation.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandMapb(args string) error {
	if pokedexLocation.Previous == nil {
	  return errors.New("No more")
	}
	err := internal.GetRequest(*pokedexLocation.Previous, &pokedexLocation)
	if err != nil {
		return err
	}
	for _, v := range pokedexLocation.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandExplore(args string) error {
	loc := strings.Split(args, " ")
	if len(loc) != 2 {
		return errors.New("Not enough arguments passed")
	}
	fmt.Println("Exploring", loc[1]+"...")
	url := fmt.Sprintf("%s/%s/%s", internal.BASE_URL, "location-area", loc[1])
	err := internal.GetRequest(url, &pokedexLocationArea)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, v := range pokedexLocationArea.PokemonEncounters {
		fmt.Println(" - ", v.Pokemon.Name)
	}
	return nil
}

func commandCatch(args string) error {
	poke := strings.Split(args, " ")
	if len(poke) != 2 {
		return errors.New("Not enough arguments passed")
	}
	url := fmt.Sprintf("%s/%s/%s", internal.BASE_URL, "pokemon", poke[1])
	fmt.Printf("Throwing a Pokeball at %s...\n", poke[1])
	err := internal.GetRequest(url, &pokedexCatch)
	if err != nil {
		return err
	}
	exp := pokedexCatch.BaseExperience
	guess := rand.Int() % exp + 10
	if guess == exp {
		UserPokemon[poke[1]] = pokedexCatch
		fmt.Println(poke[1], "was caught!")
	} else {
		fmt.Println(poke[1], "escaped!")
	}
	return nil
}

func commandInspect(args string) error {
	poke := strings.Split(args, " ")
	if len(poke) != 2 {
		return errors.New("Not enough arguments passed")
	}
	if v, ok := UserPokemon[poke[1]]; ok {
		fmt.Println("Name:", poke[1])
		fmt.Println("Height:", v.Height)
		fmt.Println("Weight:", v.Weight)

		fmt.Println("Stats:")
		for _, v := range v.Stats {
			fmt.Printf(" - %s: %d\n", v.Stat.Name, v.BaseStat)
		}
		fmt.Println("Types:")
		for _, v := range v.Types {
			fmt.Printf(" - %s\n", v.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}


func commandPokedex(args string) error {
	fmt.Println("Your Pokedex:")
	for k, _ := range UserPokemon {
		fmt.Println(" - ", k)
	}
	return nil
}
