package nfa

import (
	"log"
	"github.com/dekuu5/FiniteStateMachine/utils"
)

type NFiniteAutomata = utils.NFiniteAutomata

// ValidateNfa validates the given NFA based on the following rules:
// 1. The set of states must not be empty.
// 2. The start state must be in the set of states.
// 3. The set of input symbols must not be empty.
// 4. The set of accept states must not be empty and must be a subset of the set of states.
// 5. Each state must have transitions defined for each input symbol, and the next states must be in the set of states.
func ValidateNfa(nfa utils.NFiniteAutomata) bool {
	return validateStates(nfa) &&
		validateStartState(nfa) &&
		validateSymbols(nfa) &&
		validateAcceptStates(nfa) &&
		validateTransitions(nfa)
}

func validateStates(nfa NFiniteAutomata) bool {
	if len(nfa.States) == 0 {
		log.Println("Set of states is empty")
		return false
	}
	return true
}

func validateStartState(nfa NFiniteAutomata) bool {
	for _, state := range nfa.States {
		if state == nfa.StartState {
			return true
		}
	}
	log.Println("Start state is not in the set of states")
	return false
}

func validateSymbols(nfa NFiniteAutomata) bool {
	if len(nfa.Symbols) == 0 {
		log.Println("Set of inputs is empty")
		return false
	}
	return true
}

func validateAcceptStates(nfa NFiniteAutomata) bool {
	if len(nfa.AcceptStates) == 0 {
		log.Println("Set of accepted states is empty")
		return false
	}
	for _, acceptState := range nfa.AcceptStates {
		if !stateExists(nfa.States, acceptState) {
			log.Printf("Accepted state %s is not in the set of states", acceptState)
			return false
		}
	}
	return true
}

func validateTransitions(nfa NFiniteAutomata) bool {
	for state, transitions := range nfa.Transitions {
		if !stateExists(nfa.States, state) {
			log.Printf("State %s in transition table is not in the set of states", state)
			return false
		}
		for input, nextStates := range transitions {
			if !symbolExists(nfa.Symbols, input) {
				log.Printf("Input %s in transition table for state %s is not in the set of inputs", input, state)
				return false
			}
			for _, nextState := range nextStates {
				if !stateExists(nfa.States, nextState) {
					log.Printf("Next state %s in transition table for state %s is not in the set of states", string(nextState), state)
					return false
				}
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
