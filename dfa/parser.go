package dfa


import (
	"container/list"
	
)


func (dfaTree *DFA)ValidateString( symbols []rune) bool {
    // Initialize the queue with the input symbols
    queue := list.New()
    for _, symbol := range symbols {
        queue.PushBack(symbol)
    }

    // Start from the start state
    currentNode := dfaTree.StartState

    // Process the queue
    for queue.Len() > 0 {
        element := queue.Front()
        symbol := element.Value.(rune)
        queue.Remove(element)

        nextNode, exists := currentNode.Transitions[symbol]
        if !exists {
            return false
        }
        currentNode = nextNode
    }

    // Check if the current state is an accepting state
    return currentNode.IsAccepting
}
