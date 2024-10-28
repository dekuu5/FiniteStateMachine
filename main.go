package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"

	"github.com/dekuu5/dfa-validator/dfa"
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
	dfaJson := utils.ReadJson(*filePath)
	valid := dfa.ValidateDfa(dfaJson)
	if !valid {
        log.Fatalf("Error validating the dfa")
		os.Exit(-1)
	}
	printDfaJson(dfaJson)

	
	
	// loop to get the input string
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a string to validate using the DFA: ")
    
    // Read input until newline and trim any extra whitespace
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    // Remove the newline character from the end of the input
	symbols := []rune(strings.TrimSpace(input))

	fmt.Println(symbols)
    
    fmt.Println("You entered:", input)

	
}


func printDfaJson(dfaJson dfa.JsonDfa){
	fmt.Printf("States: %v\n", dfaJson.States)
	fmt.Printf("Symbols: %v\n", dfaJson.Symbols)
	fmt.Printf("Start State: %s\n", dfaJson.StartState)
	fmt.Printf("Accept States: %v\n", dfaJson.AcceptStates)
	fmt.Println("Transitions:")
	for state, transitions := range dfaJson.Transitions {
		fmt.Printf("  %s: %v\n", state, transitions)
	}
}