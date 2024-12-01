package dfa

/**
 * description: This file contains the implementation of the ValidateString function based on the following rules:
 * 1. The function should take a slice of runes as input.
 * 2. The function should return a boolean value.
 * 3. The function should validate the input string based on the DFA.
 */

import (
	"container/list"
)

/**
 * description: This function validates the input string based on the DFA.
 * @param symbols: A slice of runes that represents the input string
 * @return A boolean value
 */
func (dfaTree *DFA) ValidateString(symbols []rune) bool {
	// Initialize the queue with the input symbols
	queue := list.New()
	// Add the symbols to the queue
	for _, symbol := range symbols {
		queue.PushBack(symbol)
	}

	// Start from the start state
	currentNode := dfaTree.StartState

	// Process the queue
	for queue.Len() > 0 {
		// Get the first element from the queue
		element := queue.Front()
		symbol := element.Value.(rune)
		queue.Remove(element)
		// Check if the current state has a transition for the symbol
		nextNode, exists := currentNode.Transitions[symbol]
		if !exists {
			return false
		}
		// Move to the next state
		currentNode = nextNode
	}

	// Check if the current state is an accepting state
	return currentNode.IsAccepting
}
