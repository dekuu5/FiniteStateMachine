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
	//initialize the stack
	stack := NewStack()
	//initialize the queue with the input symbols
	for i := len(input) - 1; i >= 0; i-- {
		stack.Push(rune(input[i]))
	}
	// helper function to traverse the tree
	// and check if the last node is an accepting state
	var traverseTree func(currentNode *treeNode) bool
	traverseTree = func(currentNode *treeNode) bool {
		poped := false
		//base case: when the current node is an accepting state
		if stack.IsEmpty() {
			return currentNode.accepting
		}
		if currentNode.children == nil {
			return false // return false if there is no path to an accepting state

		}
		//get the current symbol
		currentSymbol := stack.Peek().(rune)
		// loop through the transitions of the current node
		for _, childNode := range currentNode.children {
			if childNode.symbol != currentSymbol && childNode.symbol != 95 {
				break
			}
			if childNode.symbol == currentSymbol {
				stack.Pop() // remove the current symbol from the queue
				poped = true
			}

			// Recursively traverse the tree for each child node
			if traverseTree(childNode) { // if the child node is an accepting state
				return true // return true and stop the traversal
			}
			continue // continue to the next child node if the current child node is not an accepting state

		}
		if poped {
			stack.Push(currentSymbol) // add the current symbol back to the queue
		}
		return false // return false if there is no path to an accepting state

	}
	// start traversing the tree from the start state
	return traverseTree(tree)
}
func (nfa *NFA) ParseTree(input string) *treeNode {
	//initialize the stack
	stack := NewStack()
	//initialize the queue with the input symbols
	for i := len(input) - 1; i >= 0; i-- {
		stack.Push(rune(input[i]))
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
		if stack.IsEmpty() {

			return node
		}
		//get the current symbol
		currentSymbol := stack.Pop().(rune)

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
		stack.Push(currentSymbol)
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
func (node *treeNode) PrintTree() {
	fmt.Print("Node Value: ", node.value)
	fmt.Print(" Accepting: ", node.accepting)
	fmt.Print(" Symbol:", string(node.symbol))
	fmt.Print(" Children: ")
	for _, child := range node.children {
		child.PrintTree()
	}
}

func (node *treeNode) PrintTreeNode() {
	fmt.Print("Node Value: ", node.value)
	fmt.Print(" Accepting: ", node.accepting)
	fmt.Print(" Symbol:", string(node.symbol))
	fmt.Print(" Children: [")
	for _, child := range node.children {
		fmt.Print(" Child Value:", child.value)
	}
	fmt.Println(" ]")
}
