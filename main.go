package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dekuu5/dfa-validator/utils"
	"github.com/dekuu5/dfa-validator/dfa"
)

func main() {
	// Define a command-line flag for the JSON file
	filePath := flag.String("file", "", "Path to the JSON file containing the DFA")
	flag.Parse()

	// Check if the file path is provided
	if *filePath == "" {
		log.Fatal("Please provide the path to the JSON file using the -file flag")
	}

	// Read the DFA from the provided JSON file
	dfaJson := utils.ReadJson(*filePath)
	valid := dfa.ValidateDfa(dfaJson)
	if !valid {
        log.Fatalf("Error validating the dfa")
		os.Exit(-1)
	}
	fmt.Printf("States: %v\n", dfaJson.States)
	fmt.Printf("Symbols: %v\n", dfaJson.Symbols)
	fmt.Printf("Start State: %s\n", dfaJson.StartState)
	fmt.Printf("Accept States: %v\n", dfaJson.AcceptStates)
	fmt.Println("Transitions:")
	for state, transitions := range dfaJson.Transitions {
		fmt.Printf("  %s: %v\n", state, transitions)
	}
}