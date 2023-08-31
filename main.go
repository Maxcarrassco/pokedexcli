package main

import (
	"fmt"
)



func main () {
  for true {
	  fmt.Print("pokedex> ")
	  var token string
	  fmt.Scanln(&token)
	  if token == "help" {
		  CommandHelp()
		  continue
	  }
	  if cmd, ok := Mapping[token]; ok {
		  cmd.CallbackFn(token)
	  } else {
		  fmt.Printf("Command %s is unknown\nTry 'help'\n", token)
	  }
  }
}
