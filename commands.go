package main


import (
	"fmt"
	"os"
	"github.com/Maxcarrassco/pokedexcli/models"
	"github.com/Maxcarrassco/pokedexcli/internal"
	"errors"
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
}


var pokedexLocation = models.PokedexLocation{}

func commandExit(args...string) error {
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


func commandMap(args...string) error {
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

func commandMapb(args...string) error {
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

// func commandExplore(args...string) error {
// }
