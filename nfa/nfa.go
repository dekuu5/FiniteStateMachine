package nfa
import (
	"fmt"
	
	"github.com/dekuu5/FiniteStateMachine/utils"
)



type StateNode struct {
    StateName   string
    Transitions map[rune][]*StateNode
    IsAccepting bool               
}



type NFA struct {
    States       []string
    Symbols     []rune
    Transitions  map[string]map[rune]string
	StartState *StateNode 
    AcceptStates []string
}

func constructNodes(jsonInput utils.NFiniteAutomata) *StateNode {
    nodes := make(map[string]*StateNode)
    for _, state := range jsonInput.States {
        nodes[state] = &StateNode{
            StateName:   state,
            Transitions: make(map[rune][]*StateNode),
            IsAccepting: false,
        }
    }

    for _, acceptState := range jsonInput.AcceptStates {
        if node, exists := nodes[acceptState]; exists {
            node.IsAccepting = true
        }
    }

    for state, transition := range jsonInput.Transitions {
        for symbol, targetStates := range transition {
            if len(symbol) == 1 {
                for _, targetState := range targetStates {
                    nodes[state].Transitions[rune(symbol[0])] = nodes[targetState]
                }
            }
        }
    }

    return nodes[jsonInput.StartState]

}
func Constructor(jsonInput utils.NFiniteAutomata) *NFA {
    transitions := make(map[string]map[rune]string)
    for state , transition := range jsonInput.Transitions {
        t := make(map[rune]string)
        for k, m := range transition {
        if len(k) == 1 {
                for _, targetState := range m {
                    t[rune(k[0])] = targetState
                }
            } else {
                fmt.Printf("Skipping key '%s' because it's not a single character\n", k)
        }
        transitions[state] = t
    }   
    }
    
    symbols := make([]rune, 0)
    for _,c := range jsonInput.Symbols{
        if len(c) == 1 {
                symbols = append(symbols, rune(c[0]))
            } else {
                fmt.Printf("Skipping key '%s' because it's not a single character\n", c)
        }
    }
    nfa := &NFA{
        States: jsonInput.States,
        Symbols: symbols,
        Transitions: transitions,
        StartState: constructNodes(jsonInput),
        AcceptStates: jsonInput.AcceptStates,
    }
    

    return nfa
    
}