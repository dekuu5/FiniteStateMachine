package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
)


type JsonDfa struct {
	States      []string                       `json:"states"`
    Symbols     []string                       `json:"symbols"`
    StartState  string                         `json:"start_state"`
    AcceptStates []string                      `json:"accept_states"`
    Transitions map[string]map[string]string   `json:"transitions"`
}


func ReadJson(fileName string) JsonDfa {
	file , err := os.Open(fileName)
	if err != nil {
        log.Fatalf("Error opening file: %v", err)
		os.Exit(-1)
	}
	defer file.Close()

	fileBytes , err := io.ReadAll(file)
	if err != nil {
        log.Fatalf("Error reading file: %v", err)
		os.Exit(-1)
	}
	var dfa JsonDfa;
	err = json.Unmarshal(fileBytes, &dfa);
	if err != nil {
        log.Fatalf("Error parsing json: %v", err)
		os.Exit(-1)
	}
	return dfa
}


// creating a function to cheak the following contions
// set of states is not empty 
// there should be a start state that in the set of states
// set of inputs are not empty 
// set of acppented states are in set of states and is not empty
// cheak the trastion table to have the same set of input and the same set of states and each states has output to the number of inputs
func ValidateDfa(dfa JsonDfa) bool {
	// Check if the set of states is not empty
	if len(dfa.States) == 0 {
		log.Println("Set of states is empty")
		return false
	}

	// Check if the start state is in the set of states
	startStateValid := false
	for _, state := range dfa.States {
		if state == dfa.StartState {
			startStateValid = true
			break
		}
	}
	if !startStateValid {
		log.Println("Start state is not in the set of states")
		return false
	}

	// Check if the set of inputs is not empty
	if len(dfa.Symbols) == 0 {
		log.Println("Set of inputs is empty")
		return false
	}

	// Check if the set of accepted states is in the set of states and is not empty
	if len(dfa.AcceptStates) == 0 {
		log.Println("Set of accepted states is empty")
		return false
	}
	for _, acceptState := range dfa.AcceptStates {
		acceptStateValid := false
		for _, state := range dfa.States {
			if state == acceptState {
				acceptStateValid = true
				break
			}
		}
		if !acceptStateValid {
			log.Printf("Accepted state %s is not in the set of states", acceptState)
			return false
		}
	}

	// Check the transition table
	for state, transitions := range dfa.Transitions {
		// Check if the state is in the set of states
		stateValid := false
		for _, s := range dfa.States {
			if s == state {
				stateValid = true
				break
			}
		}
		if !stateValid {
			log.Printf("State %s in transition table is not in the set of states", state)
			return false
		}

		// Check if the transitions have the same set of inputs
		if len(transitions) != len(dfa.Symbols) {
			log.Printf("State %s does not have transitions for all inputs", state)
			return false
		}
		for input := range transitions {
			inputValid := false
			for _, symbol := range dfa.Symbols {
				if symbol == input {
					inputValid = true
					break
				}
			}
			if !inputValid {
				log.Printf("Input %s in transition table for state %s is not in the set of inputs", input, state)
				return false
			}
		}

		// Check if each state has output to the number of inputs
		for _, nextState := range transitions {
			nextStateValid := false
			for _, s := range dfa.States {
				if s == nextState {
					nextStateValid = true
					break
				}
			}
			if !nextStateValid {
				log.Printf("Next state %s in transition table for state %s is not in the set of states", nextState, state)
				return false
			}
		}
	}

	return true
}