package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dekuu5/FiniteStateMachine/dfa"
	"github.com/dekuu5/FiniteStateMachine/nfa"
	"github.com/dekuu5/FiniteStateMachine/utils"
)

/**
 * This is the main function that reads the JSON file and processes the automaton based on the type
 * @param filePath: A string that represents the path to the JSON file
 * @param automatonType: A string that represents the type of the automaton (DFA or NFA)
 * @error: An error that indicates if there was an error reading the JSON file
 */

func main() {
	// Define command-line flags for the JSON file and type (DFA or NFA)
	filePath := flag.String("file", "", "Path to the JSON file containing the automaton")
	automatonType := flag.String("type", "dfa", "Type of the automaton (dfa or nfa)")
	flag.Parse()

	// Check if the file path is provided
	if *filePath == "" {
		log.Fatal("Please provide the path to the JSON file using the -file flag")
	}

	// Read the automaton from the provided JSON file

	// Validate and process based on the automaton type
	switch strings.ToLower(*automatonType) {
	case "dfa":
		fmt.Println("DFA")
		automatonJson := utils.ReadJson(*filePath)

		if valid := dfa.ValidateDfa(automatonJson); !valid {
			log.Fatalf("Error validating the DFA")
			os.Exit(-1)
		}
		// printDfaJson(automatonJson)

		processDfa(automatonJson)
	case "nfa":
		fmt.Println("NFA")
		automatonJson := utils.ReadJsonNfa(*filePath)
		if valid := nfa.ValidateNfa(automatonJson); !valid {
			log.Fatalf("Error validating the NFA")
			os.Exit(-1)
		}
		// nfaTree := nfa.Constructor(automatonJson)

		// printNfa(*nfaTree)
		processNfa(automatonJson)
	default:
		log.Fatalf("Unknown automaton type: %s", *automatonType)
		os.Exit(-1)
	}
}

func processDfa(dfaJson dfa.FiniteAutomata) {
	// Loop to get the input string
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

	dfaTree := dfa.Constructor(dfaJson)

	if valid := dfaTree.ValidateString(symbols); valid {
		fmt.Printf("String %s is accepted\n", input)
	} else {
		fmt.Printf("String %s is rejected\n", input)
	}
}

func processNfa(nfaJson nfa.NFiniteAutomata) {
	nfaTree := nfa.Constructor(nfaJson)

	nfaTree.PrintNFA()

	// Loop to get the input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string to validate using the NFA: ")

	// Read input until newline and trim any extra whitespace
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove the newline character from the end of the input
	symbols := strings.TrimSpace(input)
	//check if the symbols is in language
	if !nfaTree.IsInputStringValid(symbols) {
		fmt.Println("The input string is valid")
		return
	}

	if valid := nfaTree.ValidateString(symbols); valid {
		fmt.Printf("String %s is accepted\n", symbols)
	} else {
		fmt.Printf("String %s is rejected\n", symbols)
	}
}

func printDfa(dfaJson dfa.DFA) {
	fmt.Printf("States: %v\n", dfaJson.States)
	fmt.Printf("Symbols: %v\n", dfaJson.Symbols)
	fmt.Printf("Start State: %v\n", dfaJson.StartState)
	fmt.Printf("Accept States: %v\n", dfaJson.AcceptStates)
	fmt.Println("Transitions:")
	for state, transitions := range dfaJson.Transitions {
		fmt.Printf("  %s: %v\n", state, transitions)
	}
}
