package nfa

import (
	"container/list"
	"fmt"
)

func (nfa *NFA) ValidateString(input string) bool {
	tree := nfa.ParseTree(input)
	if tree == nil {
		return false
	}
	queue := list.New()

	//initialize the queue with the input symbols
	for _, symbol := range input {
		queue.PushBack(rune(symbol))
	}
	// helper function to traverse the tree
	// and check if the last node is an accepting state
	var traverseTree func(currentNode *StateNode, PreviousSymbol rune) bool
	traverseTree = func(currentNode *StateNode, PreviousSymbol rune) bool {

		if queue.Len() == 0 {
			if currentNode.IsAccepting {
				return true
			}
			return false

		}
		//get the current symbol
		currentSymbol := queue.Front().Value.(rune)
		queue.Remove(queue.Front())
		// check if the current symbol has no next node in the tree
		// the branch is ended and the input string is still not finished
		if currentNode.Transitions[currentSymbol] == nil {
			return false
		}

		for _, childNode := range currentNode.Transitions[currentSymbol] {
			return traverseTree(childNode, currentSymbol)
		}
		// handle epsilon transitions
		for _, childNode := range currentNode.Transitions['ε'] {
			return traverseTree(childNode, PreviousSymbol)
		}
		return false
	}
	return traverseTree(tree, 'ε')
}
func (nfa *NFA) ParseTree(input string) *StateNode {
	queue := list.New()

	//initialize the queue with the input symbols
	for _, symbol := range input {
		queue.PushBack(rune(symbol))
	}

	// Recursive helper function to build the tree

	var buildTree func(currentState *StateNode) *StateNode
	buildTree = func(currentState *StateNode) *StateNode {

		//base case: when the current symbol is queue is empty
		if queue.Len() == 0 {
			return nil
		}
		//get the current symbol
		currentSymbol := queue.Front().Value.(rune)
		queue.Remove(queue.Front())

		//base case: when the current state is nil or there is no transition for the current symbol
		if currentState == nil || nfa.Transitions[currentState.StateName][currentSymbol] == nil {
			return nil
		}

		// Create a new node for the current state
		node := &StateNode{
			StateName:   currentState.StateName,
			Transitions: make(map[rune][]*StateNode),
			IsAccepting: currentState.IsAccepting,
		}

		// Get the slice of possible next states for the current symbol
		for _, nextState := range nfa.Transitions[currentState.StateName][currentSymbol] {
			childNode := buildTree(nfa.getNode(nextState))
			node.Transitions[currentSymbol] = append(node.Transitions[currentSymbol], childNode)
		}
		// handle epsilon transitions
		for _, nextState := range nfa.Transitions[currentState.StateName]['ε'] {
			childNode := buildTree(nfa.getNode(nextState))
			node.Transitions['ε'] = append(node.Transitions['ε'], childNode)
		}
		// Return the node
		return node
	}

	// Start building the tree from the start state of the NFA
	return buildTree(nfa.StartState)

}

// getNode returns the node corresponding to the state name
func (nfa *NFA) getNode(stateName string) *StateNode {
	for _, state := range nfa.States {
		if state == stateName {
			return nfa.StartState
		}
	}
	return nil
}

func (nfa *NFA) IsInputStringValid(input string) bool {
	for _, char := range input {
		if !containsSymbol(nfa.Symbols, char) {
			fmt.Printf("Invalid symbol '%c' in input string\n", char)
			return false
		}
	}
	return true
}

// containsSymbol checks if a symbol is in the list of symbols
func containsSymbol(symbols []rune, char rune) bool {
	for _, symbol := range symbols {
		if symbol == char {
			return true
		}
	}
	return false
}

// print the tree for debugging purposes with tree format
func (nfa *NFA) printTree(node *StateNode, level rune) {
	if node == nil {
		return
	}
	fmt.Printf("%s:%s\n", string(level), node.StateName)
	for symbol, children := range node.Transitions {
		fmt.Printf("%s -> %s\n", string(level), string(symbol))
		for _, child := range children {
			nfa.printTree(child, level+1)
		}
	}

}
