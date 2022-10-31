package main

import (
	"fmt"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	whoToGreet := "World"
	who := githubactions.GetInput("who-to-greet")
  	if who == "" {
    	githubactions.Fatalf("missing 'who-to-greet'")
  	} else {
		whoToGreet = who
	}
	fmt.Println(greet(whoToGreet))
	githubactions.SetOutput("greeting", greet(whoToGreet))
}

func greet(who string) string {
	return "Hello" + " " + who
}
