package nfa

/**
 * description: the file contains the functions that validate the given NFA based on the following rules:
 * 1. The set of states must not be empty.
 * 2. The start state must be in the set of states.
 * 3. The set of input symbols must not be empty.
 * 4. The set of accept states must not be empty and must be a subset of the set of states.
 * 5. Each state may have transitions for each input symbol, and the next states must be in the set of states or no transitions at all.
 */

import (
	"log"

	"github.com/dekuu5/FiniteStateMachine/utils"
)

type NFiniteAutomata = utils.NFiniteAutomata

/**
 * This function validates the given NFA based on mentioned rules above
 * @param nfa: A NFiniteAutomata struct that represents the NFA
 * @return A boolean that indicates if the NFA is valid
 */
func ValidateNfa(nfa utils.NFiniteAutomata) bool {
	return validateStates(nfa) && // check if the set of states is not empty
		validateStartState(nfa) && // check if the start state is in the set of states
		validateSymbols(nfa) && // check if the set of input symbols is not empty
		validateAcceptStates(nfa) && // check if the set of accept states is not empty and is a subset of the set of states
		validateTransitions(nfa) // check if the transitions are valid based on the set of states and input symbols
}

/**
 * This function validates the set of states of the NFA
 * @param nfa: A NFiniteAutomata struct that represents the NFA
 * @error: If the set of states is empty
 * @return A boolean that indicates if the set of states is valid
 */
func validateStates(nfa NFiniteAutomata) bool {
	if len(nfa.States) == 0 {
		log.Println("Set of states is empty")
		return false
	}
	return true
}

/**
 * This function validates the start state of the NFA
 * @param nfa: A NFiniteAutomata struct that represents the NFA
 * @error: If the start state is not in the set of states
 * @return A boolean that indicates if the start state is valid
 */
func validateStartState(nfa NFiniteAutomata) bool {
	for _, state := range nfa.States {
		if state == nfa.StartState {
			return true
		}
	}
	log.Println("Start state is not in the set of states")
	return false
}

/**
 * This function validates the set of input symbols of the NFA
 * @param nfa: A NFiniteAutomata struct that represents the NFA
 * @error: If the set of input symbols is empty
 * @return A boolean that indicates if the set of input symbols is valid
 */
func validateSymbols(nfa NFiniteAutomata) bool {
	if len(nfa.Symbols) == 0 {
		log.Println("Set of inputs is empty")
		return false
	}
	return true
}

/**
 * This function validates the set of accept states of the NFA
 * @param nfa: A NFiniteAutomata struct that represents the NFA
 * @error: If the set of accept states is empty or not a subset of the set of states
 * @return A boolean that indicates if the set of accept states is valid
 */

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

/**
 * This function validates the transitions of the NFA
 * @param nfa: A NFiniteAutomata struct that represents the NFA
 * @error: If the transitions are not valid based on the set of states and input symbols
 * @return A boolean that indicates if the transitions are valid
 */
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

/**
 * description: Check if the state exists in the set of states
 * @param states: A slice of strings that represents the set of states
 * @param state: A string that represents the state
 * @return A boolean that indicates if the state exists in the set of states
 */
func stateExists(states []string, state string) bool {
	for _, s := range states {
		if s == state { // if the state exists in the set of states
			return true
		}
	}
	return false
}

/**
 * description: Check if the input string is valid based on the set of input symbols
 * @param input: A string that represents the input string
 * @return A boolean that indicates if the input string is valid
 */
func symbolExists(symbols []string, symbol string) bool {
	for _, s := range symbols {
		if s == symbol {
			return true
		}
	}
	return false
}
