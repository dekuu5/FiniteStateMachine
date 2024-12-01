package dfa

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

func constructNodes(jsonInput utils.FiniteAutomata) *StateNode {
	nodes := make(map[string]*StateNode)
	for _, state := range jsonInput.States {
		nodes[state] = &StateNode{
			StateName:   state,
			Transitions: make(map[rune]*StateNode),
			IsAccepting: false,
		}
	}

	for _, acceptState := range jsonInput.AcceptStates {
		if node, exists := nodes[acceptState]; exists {
			node.IsAccepting = true
		}
	}

	for state, transition := range jsonInput.Transitions {
		for symbol, targetState := range transition {
			if len(symbol) == 1 {
				nodes[state].Transitions[rune(symbol[0])] = nodes[targetState]
			}
		}
	}

	return nodes[jsonInput.StartState]

}
func Constructor(jsonInput utils.FiniteAutomata) *DFA {
	transitions := make(map[string]map[rune]string)
	for state, transition := range jsonInput.Transitions {
		t := make(map[rune]string)
		for k, m := range transition {
			if len(k) == 1 {
				t[rune(k[0])] = m
			} else {
				fmt.Printf("Skipping key '%s' because it's not a single character\n", k)
			}
			transitions[state] = t
		}
	}

	symbols := make([]rune, 0)
	for _, c := range jsonInput.Symbols {
		if len(c) == 1 {
			symbols = append(symbols, rune(c[0]))
		} else {
			fmt.Printf("Skipping key '%s' because it's not a single character\n", c)
		}
	}
	dfa := &DFA{
		States:       jsonInput.States,
		Symbols:      symbols,
		Transitions:  transitions,
		StartState:   constructNodes(jsonInput),
		AcceptStates: jsonInput.AcceptStates,
	}

	return dfa

}
func (dfa *DFA) PrintNFA() {
	fmt.Println("States:", dfa.States)
	fmt.Print("Symbols: [ ")
	for _, symbol := range dfa.Symbols {
		fmt.Print(string(symbol), " ")
	}
	fmt.Println(" ]")
	fmt.Println("Transitions:")
	for _, state := range dfa.States {
		fmt.Print(state)
		fmt.Print(" -> [")
		for symbol, transition := range dfa.Transitions[state] {
			fmt.Print(" ", string(symbol), " : ", transition)
		}
		fmt.Println(" ]")
	}
	fmt.Println("StartState:", dfa.StartState.StateName)
	fmt.Println("AcceptStates:", dfa.AcceptStates)

}
