package utils

/**
 * description: the file contains the implementation of the non determinstic FiniteAutomata struct
 */
import (
	"encoding/json"
	"io"
	"log"
	"os"
)

/**
 * This is the struct that represents a NFA
 * It has the following fields:
 * States: A slice of strings that represents the states of the NFA
 * Symbols: A slice of strings that represents the symbols of the NFA
 * StartState: A string that represents the start state of the NFA
 * AcceptStates: A slice of strings that represents the accepting states of the NFA
 * Transitions: A map of strings to a map of strings to a slice of strings
 */
type NFiniteAutomata struct {
	States       []string                       `json:"states"`
	Symbols      []string                       `json:"symbols"`
	StartState   string                         `json:"start_state"`
	AcceptStates []string                       `json:"accept_states"`
	Transitions  map[string]map[string][]string `json:"transitions"`
}

/**
* description: This function parse a json file and returns a NFiniteAutomata struct
* @param fileName: A string that represents the name of the json file
* @error: if there is an error opening the file, reading the file, or parsing the json
* @return: A NFiniteAutomata struct that represents the NFA
 */
func ReadJsonNfa(fileName string) NFiniteAutomata {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
		os.Exit(-1)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		os.Exit(-1)
	}
	var finiteAutomata NFiniteAutomata
	err = json.Unmarshal(fileBytes, &finiteAutomata)
	if err != nil {
		log.Fatalf("Error parsing json: %v", err)
		os.Exit(-1)
	}
	return finiteAutomata
}
