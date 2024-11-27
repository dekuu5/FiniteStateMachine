package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type FiniteAutomata struct {
	States       []string                     `json:"states"`
	Symbols      []string                     `json:"symbols"`
	StartState   string                       `json:"start_state"`
	AcceptStates []string                     `json:"accept_states"`
	Transitions  map[string]map[string]string `json:"transitions"`
}

func ReadJson(fileName string) FiniteAutomata {
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
	var finiteAutomata FiniteAutomata
	err = json.Unmarshal(fileBytes, &finiteAutomata)
	if err != nil {
		log.Fatalf("Error parsing json: %v", err)
		os.Exit(-1)
	}
	return finiteAutomata
}
