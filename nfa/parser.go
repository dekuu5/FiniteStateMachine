package nfa

import (
	"fmt"
)

type treeNode struct {
	value     string
	children  []*treeNode
	accepting bool
	symbol    rune
}

func (nfa *NFA) ValidateString(input string) bool {
	tree := nfa.ParseTree(input)
	//check if the tree is nil
	if tree == nil {
		return false // the is no path in the tree
	}
	//initialize the queue
	queue := NewQueue()

	//initialize the queue with the input symbols
	for _, symbol := range input {
		queue.Enqueue(rune(symbol))
	}
	// helper function to traverse the tree
	// and check if the last node is an accepting state
	var traverseTree func(currentNode *treeNode) bool
	traverseTree = func(currentNode *treeNode) bool {
		fmt.Println("currentNode", currentNode)

		//base case: when the current node is an accepting state
		if queue.IsEmpty() {
			return currentNode.accepting
		}
		if currentNode.children == nil {
			return false // return false if there is no path to an accepting state

		}
		fmt.Println("queue", queue)
		//get the current symbol
		currentSymbol := queue.Dequeue().(rune)

		fmt.Println("currentSymbol", string(currentSymbol))
		// loop through the transitions of the current node
		for _, childNode := range currentNode.children {
			if childNode.symbol == 95 {
				fmt.Println("epsilon transition")
				queue.Enqueue(currentSymbol)
			}

			// Recursively traverse the tree for each child node
			if traverseTree(childNode) { // if the child node is an accepting state
				return true // return true and stop the traversal
			}
			continue // continue to the next child node if the current child node is not an accepting state

		}

		queue.Enqueue(currentSymbol) // add the current symbol back to the queue

		return false // return false if there is no path to an accepting state

	}
	// start traversing the tree from the start state
	return traverseTree(tree)
}
func (nfa *NFA) ParseTree(input string) *treeNode {
	queue := NewQueue()

	//initialize the queue with the input symbols
	for _, symbol := range input {
		queue.Enqueue(rune(symbol))
	}

	// Recursive helper function to build the tree

	var buildTree func(currentState *StateNode) *treeNode
	buildTree = func(currentState *StateNode) *treeNode {
		// Create a new node for the current state
		node := &treeNode{
			value:     currentState.StateName,
			children:  nil, // children will be set later
			accepting: currentState.IsAccepting,
			symbol:    0,
		}
		if len(currentState.Transitions['_']) > 0 {
			//handle if the current node has epsilon transitions
			for _, nextnode := range currentState.Transitions['_'] {
				// Recursively build the tree for each possible next state
				childNode := buildTree(nextnode)
				// Append the child node to the current node
				if childNode != nil {
					node.children = append(node.children, childNode)
					childNode.symbol = '_'
				}
			}
		}
		//base case: when the current symbol is queue is empty
		if queue.IsEmpty() {
			return node
		}
		//get the current symbol
		currentSymbol := queue.Dequeue().(rune)

		if len(currentState.Transitions[currentSymbol]) == 0 {
			return node
		}

		// Get the slice of possible next states for the current symbol
		for _, nextnode := range currentState.Transitions[currentSymbol] {
			// Recursively build the tree for each possible next state
			childNode := buildTree(nextnode)
			// Append the child node to the current node
			if childNode != nil {
				node.children = append(node.children, childNode)
				childNode.symbol = currentSymbol
			}
		}
		queue.Enqueue(currentSymbol)
		return node
	}

	// Start building the tree from the start state of the NFA
	return buildTree(nfa.StartState)

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
