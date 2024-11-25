package dfa

import (
	"log"
	"github.com/dekuu5/FiniteStateMachine/utils"
)

type FiniteAutomata = utils.FiniteAutomata




func ValidateDfa(dfa FiniteAutomata) bool {
	return validateStates(dfa) &&
		validateStartState(dfa) &&
		validateSymbols(dfa) &&
		validateAcceptStates(dfa) &&
		validateTransitions(dfa)
}


func validateStates(dfa FiniteAutomata) bool {
	if len(dfa.States) == 0 {
		log.Println("Set of states is empty")
		return false
	}
	return true
}

func validateStartState(dfa FiniteAutomata) bool {
	for _, state := range dfa.States {
		if state == dfa.StartState {
			return true
		}
	}
	log.Println("Start state is not in the set of states")
	return false
}

func validateSymbols(dfa FiniteAutomata) bool {
	if len(dfa.Symbols) == 0 {
		log.Println("Set of inputs is empty")
		return false
	}
	return true
}

func validateAcceptStates(dfa FiniteAutomata) bool {
	if len(dfa.AcceptStates) == 0 {
		log.Println("Set of accepted states is empty")
		return false
	}
	for _, acceptState := range dfa.AcceptStates {
		if !stateExists(dfa.States, acceptState) {
			log.Printf("Accepted state %s is not in the set of states", acceptState)
			return false
		}
	}
	return true
}

func validateTransitions(dfa FiniteAutomata) bool {
	for state, transitions := range dfa.Transitions {
		if !stateExists(dfa.States, state) {
			log.Printf("State %s in transition table is not in the set of states", state)
			return false
		}
		if len(transitions) != len(dfa.Symbols) {
			log.Printf("State %s does not have transitions for all inputs", state)
			return false
		}
		for input, nextState := range transitions {
			if !symbolExists(dfa.Symbols, input) {
				log.Printf("Input %s in transition table for state %s is not in the set of inputs", input, state)
				return false
			}
			if !stateExists(dfa.States, nextState) {
				log.Printf("Next state %s in transition table for state %s is not in the set of states", nextState, state)
				return false
			}
		}
	}
	return true
}

func stateExists(states []string, state string) bool {
	for _, s := range states {
		if s == state {
			return true
		}
	}
	return false
}

func symbolExists(symbols []string, symbol string) bool {
	for _, s := range symbols {
		if s == symbol {
			return true
		}
	}
	return false
}
