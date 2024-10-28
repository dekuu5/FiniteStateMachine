package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dekuu5/dfa-validator/utils"
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
	dfa := utils.ReadJson(*filePath)
	valid := utils.ValidateDfa(dfa)
	if !valid {
        log.Fatalf("Error validating the dfa")
		os.Exit(-1)
	}
	fmt.Printf("States: %v\n", dfa.States)
	fmt.Printf("Symbols: %v\n", dfa.Symbols)
	fmt.Printf("Start State: %s\n", dfa.StartState)
	fmt.Printf("Accept States: %v\n", dfa.AcceptStates)
	fmt.Println("Transitions:")
	for state, transitions := range dfa.Transitions {
		fmt.Printf("  %s: %v\n", state, transitions)
	}
}