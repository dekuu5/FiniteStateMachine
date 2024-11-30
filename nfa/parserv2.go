package nfa

import (
	"fmt"
)

func (nfa *NFA) ValidateStringDac(input []rune) bool {
	//queue := list.New()
	//
	//// Initialize the queue with the input symbols
	//for _, symbol := range input {
	//	queue.PushBack(rune(symbol))
	//}
	queue := NewQueue()

	for _, symbol := range input {
		queue.Enqueue(symbol)
	}
	return parserDac(nfa.StartState, queue)
}

//func parserDac(startState *StateNode, chars list.List) bool {
//    type stateWithChars struct {
//        state *StateNode
//        chars list.List
//    }
//
//    stack := list.New()
//    stack.PushBack(stateWithChars{startState, chars})
//
//    for stack.Len() > 0 {
//		fmt.Println(stack)
//        element := stack.Back()
//        stack.Remove(element)
//
//        current := element.Value.(stateWithChars)
//
//        currentState := current.state
//        currentChars := current.chars
//
//        if currentChars.Len() == 0 {
//            if currentState.IsAccepting {
//                return true
//            }
//            continue
//        }
//
//        nextChar := currentChars.Remove(currentChars.Front()).(rune)
//        if nextStates, ok := currentState.Transitions[nextChar]; ok {
//            for _, nextState := range nextStates {
//                newChars := list.New()
//                for e := currentChars.Front(); e != nil; e = e.Next() {
//                    newChars.PushBack(e.Value)
//                }
//                stack.PushBack(stateWithChars{nextState, *newChars})
//            }
//        }
//
//        // Check for epsilon transitions (empty transitions)
//        if epsilonTransitions, ok := currentState.Transitions['_']; ok {
//            for _, nextState := range epsilonTransitions {
//                newChars := list.New()
//                for e := currentChars.Front(); e != nil; e = e.Next() {
//                    newChars.PushBack(e.Value)
//                }
//                stack.PushBack(stateWithChars{nextState, *newChars})
//            }
//        }
//    }
//
//    return false
//}

func parserDac(startState *StateNode, chars *Queue) bool {
	if chars.Size() == 0 {
		return startState.IsAccepting
	}

	nextChar := chars.Dequeue().(rune)
	fmt.Println(startState.StateName, nextChar, chars.Size())

	// Check for transitions with the current character
	if nextStates, ok := startState.Transitions[nextChar]; ok {
		for _, nextState := range nextStates {
			newChars := chars.Copy()
			if parserDac(nextState, newChars) {
				return true
			}
		}
	}

	// Check for epsilon transitions (empty transitions)
	if epsilonTransitions, ok := startState.Transitions['_']; ok {
		for _, nextState := range epsilonTransitions {
			newChars := chars.Copy()
			if parserDac(nextState, newChars) {
				return true
			}
		}
	}

	return false
}
