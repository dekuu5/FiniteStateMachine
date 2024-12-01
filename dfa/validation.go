package dfa

/**
 * description: validate the DFA based on the set of states, start state, input symbols, accepted states, and transitions based the following rules:
 * 1. The set of states must not be empty.
 * 2. The start state must be in the set of states.
 * 3. The set of input symbols must not be empty.
 * 4. The set of accepted states must not be empty and must be a subset of the set of states.
 * 5. Each state must have one transition for each input symbol, and the next states must be in the set of states.
 */

import (
	"log"

	"github.com/dekuu5/FiniteStateMachine/utils"
)

// FiniteAutomata is a type alias for the FiniteAutomata struct in the utils package
type FiniteAutomata = utils.FiniteAutomata

/**
 * description: validate the given DFA based on the mentioned rules above
 * @param dfa: A FiniteAutomata struct that represents the DFA
 * @return: A boolean that indicates if the DFA is valid
 */
func ValidateDfa(dfa FiniteAutomata) bool {
	//combine all the validation functions and return the result
	return validateStates(dfa) &&
		validateStartState(dfa) &&
		validateSymbols(dfa) &&
		validateAcceptStates(dfa) &&
		validateTransitions(dfa)
}

/**
 * description: validate the set of states of the DFA
 * @param dfa: A FiniteAutomata struct that represents the DFA
 * @error: if the set of states is empty
 * @return: A boolean that indicates if the set of states is valid
 */
func validateStates(dfa FiniteAutomata) bool {
	//check if the set of states is not empty
	if len(dfa.States) == 0 {
		log.Println("Set of states is empty")
		return false
	}
	return true
}

/**
 * description: validate the start state of the DFA
 * @param dfa: A FiniteAutomata struct that represents the DFA
 * @error: if the start state is not in the set of states
 * @return: A boolean that indicates if the start state is valid
 */

func validateStartState(dfa FiniteAutomata) bool {
	//check if the start state is in the set of states
	for _, state := range dfa.States {
		if state == dfa.StartState {
			return true
		}
	}
	//if the start state is not in the set of states
	log.Println("Start state is not in the set of states")
	return false
}

/**
 * description: validate the set of input symbols of the DFA
 * @param dfa: A FiniteAutomata struct that represents the DFA
 * @error: if the set of input symbols is empty
 * @return: A boolean that indicates if the set of input symbols is valid
 */

func validateSymbols(dfa FiniteAutomata) bool {
	//check if the set of input symbols is not empty
	if len(dfa.Symbols) == 0 {
		log.Println("Set of inputs is empty")
		return false
	}
	return true
}

/**
 * description: validate the accepted states of the DFA
 * @param dfa: A FiniteAutomata struct that represents the DFA
 * @error: if the set of accepted states is empty or not a subset of the set of states
 * @return: A boolean that indicates if the accepted states are valid
 */
func validateAcceptStates(dfa FiniteAutomata) bool {
	//check if the set of accepted states is not empty
	if len(dfa.AcceptStates) == 0 {
		log.Println("Set of accepted states is empty")
		return false
	}
	//check if the accepted states are a subset of the set of states
	for _, acceptState := range dfa.AcceptStates {
		//check if the accepted state is in the set of states
		if !stateExists(dfa.States, acceptState) {
			log.Printf("Accepted state %s is not in the set of states", acceptState)
			return false
		}
	}
	return true
}

/**
 * description: validate the transitions of the DFA based on the set of states and input symbols
 * @param dfa: a DFA struct that represents the DFA
 * @error: if a state in the transition table is not in the set of states, an input in the transition table is not in the set of inputs,
 * a state has more than one next state for an input, or a next state in the transition table is not in the set of states
 * @return: a boolean value
 */
func validateTransitions(dfa FiniteAutomata) bool {
	for state, transitions := range dfa.Transitions {
		//check if a state in the transition table is in the set of states
		if !stateExists(dfa.States, state) {
			log.Printf("State %s in transition table is not in the set of states", state)
			return false
		}
		//check if a state has transitions for all inputs
		if len(transitions) != len(dfa.Symbols) {
			log.Printf("State %s does not have transitions for all inputs", state)
			return false
		}
		for input, nextState := range transitions {
			//check if an input in the transition table is in the set of inputs
			if !symbolExists(dfa.Symbols, input) {
				log.Printf("Input %s in transition table for state %s is not in the set of inputs", input, state)
				return false
			}
			//check if a state has more than one next state for an input
			if len(nextState) > 1 {
				log.Printf("State %s has more than one next state for input %s", state, input)
			}
			//check if the next state is in the set of states
			if !stateExists(dfa.States, nextState) {
				log.Printf("Next state %s in transition table for state %s is not in the set of states", nextState, state)
				return false
			}
		}
	}
	return true
}

/**
 * description: check if a state exists in the list of states
 * @param states: a list of states, type []string
 * @param state: a state to check, type string
 * @return: a boolean value
 */

func stateExists(states []string, state string) bool {
	// loop through the states and check if the state exists in the set of states
	for _, s := range states {
		if s == state {
			return true // if the state exists in the set of states
		}
	}
	return false
}

/**
 * description: check if a symbol exists in the list of symbols
 * @param symbols: a list of symbols, type []string
 * @param symbol: a symbol to check, type string
 * @return: a boolean value
 */

func symbolExists(symbols []string, symbol string) bool {
	// loop through the symbols and check if the symbol exists in the set of symbols
	for _, s := range symbols {
		if s == symbol {
			return true // if the symbol exists in the set of symbols
		}
	}
	return false
}
