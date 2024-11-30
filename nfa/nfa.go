package nfa

import (
	"fmt"

	"github.com/dekuu5/FiniteStateMachine/utils"
)

/**
 * This is the struct that represents a state node in the NFA
 * It has the following fields:
 * StateName: The name of the state
 * Transitions: A map of runes to a slice of pointers to StateNodes
 * IsAccepting: A boolean that indicates if the state is an accepting state
 */

type StateNode struct {
	StateName   string
	Transitions map[rune][]*StateNode
	IsAccepting bool
}

/**
 * This is the struct that represents a NFA
 * It has the following fields:
 * States: A slice of strings that represents the states of the NFA
 * Symbols: A slice of runes that represents the symbols of the NFA
 * Transitions: A map of strings to a map of runes to a slice of strings
 * StartState: A pointer to the start state of the NFA
 * AcceptStates: A slice of strings that represents the accepting states of the NFA
 */
type NFA struct {
	States       []string
	Symbols      []rune
	Transitions  map[string]map[rune][]string
	StartState   *StateNode
	AcceptStates []string
}

/**
 * This function constructs the nodes of the NFA
 * @param jsonInput: A NFiniteAutomata struct that represents the NFA
 * @return A pointer to the start state of the NFA
 */
func constructNodes(jsonInput utils.NFiniteAutomata) *StateNode {
	nodes := make(map[string]*StateNode) // map of state name to StateNode

	for _, state := range jsonInput.States { // loop through the states and create a StateNode for each state
		nodes[state] = &StateNode{
			StateName:   state,                       // set the state name
			Transitions: make(map[rune][]*StateNode), // create a map of runes to a slice of pointers to StateNodes
			IsAccepting: false,                       // set the IsAccepting field to false
		}
	}

	for _, acceptState := range jsonInput.AcceptStates { // loop through the accepting states and set the IsAccepting field to true
		if node, exists := nodes[acceptState]; exists { //check if the state exists in the nodes map
			node.IsAccepting = true // set the IsAccepting field to true
		}

	}

	for state, transition := range jsonInput.Transitions { // loop through the transitions and set the transitions of each state
		for symbol, targetStates := range transition { // loop through the target states of each symbol
			if len(symbol) == 1 {
				for _, targetState := range targetStates {
					// append the target state to the transitions of the state
					nodes[state].Transitions[rune(symbol[0])] = append(nodes[state].Transitions[rune(symbol[0])], nodes[targetState])
				}
			}
		}
	}

	return nodes[jsonInput.StartState]

}

/**
 * This is the constructor function for the NFA struct
 * It constructs the NFA struct from a NFiniteAutomata struct
 * @param jsonInput: A NFiniteAutomata struct that represents the NFA
 * @return A pointer to the constructed NFA struct
 */
func Constructor(jsonInput utils.NFiniteAutomata) *NFA {

	
	// create a map of strings to a map of runes to a slice of strings
	transitions := make(map[string]map[rune][]string)
	// loop through the transitions and set the transitions of each state
	for state, transition := range jsonInput.Transitions {
		// create a map of runes to a slice of strings
		t := make(map[rune][]string)
		// loop through the transitions of each state
		// fmt.Println(transition)
		for k, m := range transition {
			fmt.Println(k,[]rune(k))
			if len(k) == 1 {
				// loop through the target states of each symbol
				for _, targetState := range m {
					t[rune(k[0])] = append(t[rune(k[0])], targetState) // append the target state to the transitions of the state
				}
			}else {
				fmt.Printf("Skipping key '%s' because it's not a single character\n", k)
			}
			transitions[state] = t
		}
	}
	// create a slice of runes
	symbols := make([]rune, 0)
	for _, c := range jsonInput.Symbols {
		if len(c) == 1 {
			symbols = append(symbols, rune(c[0])) // append the rune to the symbols slice
		} else {
			fmt.Printf("Skipping key '%s' because it's not a single sqs character\n", c) // i think this is a bug //throw "Skipping key 'ε' because it's not a single character" error
		}
	}
	// create a NFA struct
	nfa := &NFA{
		States:       jsonInput.States,
		Symbols:      symbols,
		Transitions:  transitions,
		StartState:   constructNodes(jsonInput),
		AcceptStates: jsonInput.AcceptStates,
	}

	return nfa

}
