package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	suggest "github.com/picatz/suggest/core"
)

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			os.Exit(0)
		}
	}()
}

var (
	StatusNoArgs        = 1
	StatusGetErr        = 2
	StatusNoSuggestions = 3
)

func main() {

	args := os.Args[1:]

	if len(args) <= 0 {
		fmt.Println("no arguments given to get suggestions for")
		os.Exit(StatusNoArgs)
	}

	suggestions, err := suggest.Get(strings.Join(args, " "))

	if err != nil {
		fmt.Println(err)
		os.Exit(StatusGetErr)
	}

	if len(suggestions) == 0 {
		os.Exit(StatusNoSuggestions)
	}

	for _, suggestion := range suggestions {
		fmt.Println(suggestion)
	}
}
