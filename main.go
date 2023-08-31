package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main () {
  for true {
	  fmt.Print("Pokedex > ")
	  stdIn := bufio.NewScanner(os.Stdin)
	  stdIn.Split(bufio.ScanLines)
	  stdIn.Scan()
	  token := stdIn.Text()
	  if token == "help" {
		  CommandHelp()
		  continue
	  }
	  args := strings.Split(token, " ")
	  if cmd, ok := Mapping[args[0]]; ok {
		  cmd.CallbackFn(token)
	  } else {
		  fmt.Printf("Command %s is unknown\nTry 'help'\n", token)
	  }
  }
}
