package dfa

/**
 * description: This file contains the implementation of the DFA struct and its methods
 */
import (
	"fmt"

	"github.com/dekuu5/FiniteStateMachine/utils"
)

/**
 * This is the struct that represents a state node in the DFA
 * It has the following fields:
 * StateName: The name of the state
 * Transitions: A map of runes to a pointer to StateNode
 * IsAccepting: A boolean that indicates if the state is an accepting state
 */
type StateNode struct {
	StateName   string
	Transitions map[rune]*StateNode
	IsAccepting bool
}

/**
 * This is the struct that represents a DFA
 * It has the following fields:
 * States: A slice of strings that represents the states of the DFA
 * Symbols: A slice of runes that represents the symbols of the DFA
 * Transitions: A map of strings to a map of runes to a string
 * StartState: A pointer to the start state of the DFA
 * AcceptStates: A slice of strings that represents the accepting states of the DFA
 */

type DFA struct {
	States       []string
	Symbols      []rune
	Transitions  map[string]map[rune]string
	StartState   *StateNode
	AcceptStates []string
}

/**
 * description: This function constructs the nodes of the DFA
 * @param jsonInput: A FiniteAutomata struct that represents the DFA
 * @return A pointer to the start state of the DFA
 */
func constructNodes(jsonInput utils.FiniteAutomata) *StateNode {
	nodes := make(map[string]*StateNode)
	// Loop through the states and create a StateNode for each state
	for _, state := range jsonInput.States {
		nodes[state] = &StateNode{
			StateName:   state,
			Transitions: make(map[rune]*StateNode),
			IsAccepting: false,
		}
	}
	// Loop through the accepting states and set the IsAccepting field to true
	for _, acceptState := range jsonInput.AcceptStates {
		if node, exists := nodes[acceptState]; exists {
			node.IsAccepting = true
		}
	}
	// Loop through the transitions and add the transitions to the nodes
	for state, transition := range jsonInput.Transitions {
		for symbol, targetState := range transition {
			if len(symbol) == 1 {
				nodes[state].Transitions[rune(symbol[0])] = nodes[targetState]
			}
		}
	}
	// Return the start state
	return nodes[jsonInput.StartState]

}

/**
 * description: This function constructs a DFA struct based on the given JSON input
 * @param jsonInput: A FiniteAutomata struct that represents the DFA
 * @return A pointer to the constructed DFA struct
 */
func Constructor(jsonInput utils.FiniteAutomata) *DFA {
	// Construct the transitions map
	transitions := make(map[string]map[rune]string)
	for state, transition := range jsonInput.Transitions {
		t := make(map[rune]string)
		// Loop through the transitions for each state
		for k, m := range transition {
			if len(k) == 1 {
				// Convert the key to a rune and add the transition to the map
				t[rune(k[0])] = m
			} else {
				fmt.Printf("Skipping key '%s' because it's not a single character\n", k)
			}
			// Add the transitions map to the dfa transitions map
			transitions[state] = t
		}
	}

	// Construct the symbols slice
	symbols := make([]rune, 0)
	for _, c := range jsonInput.Symbols {
		if len(c) == 1 {
			symbols = append(symbols, rune(c[0]))
		} else {
			fmt.Printf("Skipping key '%s' because it's not a single character\n", c)
		}
	}
	// Construct the DFA struct with the checked input
	dfa := &DFA{
		States:       jsonInput.States,
		Symbols:      symbols,
		Transitions:  transitions,
		StartState:   constructNodes(jsonInput),
		AcceptStates: jsonInput.AcceptStates,
	}
	// Return the constructed DFA struct
	return dfa

}

/**
 * description: Print the DFA struct
 * @param dfa: a DFA struct that represents the DFA
 */
func (dfa *DFA) PrintDFA() {
	// Print the DFA struct
	//  Print the states
	fmt.Println("States:", dfa.States)
	// Print the symbols
	fmt.Print("Symbols: [ ")
	for _, symbol := range dfa.Symbols {
		fmt.Print(string(symbol), " ")
	}
	fmt.Println(" ]")
	// Print the transitions
	fmt.Println("Transitions:")
	for _, state := range dfa.States {
		fmt.Print(state)
		fmt.Print(" -> [")
		// loop through the transitions for each state
		for symbol, transition := range dfa.Transitions[state] {

			fmt.Print(" ", string(symbol), " : ", transition)
		}
		fmt.Println(" ]")
	}
	// Print the start state and the accepting states
	fmt.Println("StartState:", dfa.StartState.StateName)
	fmt.Println("AcceptStates:", dfa.AcceptStates)

}
