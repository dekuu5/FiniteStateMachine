package dfa

// import "github.com/dekuu5/dfa-validator/utils"



type StateNode struct {
    StateName   string
    Transitions map[rune]*StateNode
    IsAccepting bool               
}



type DFA struct {
    States       []string
    Alphabet     []rune
    Transitions  map[string]map[rune]string
	StartState *StateNode 
    AcceptStates []string
}


// FIRST MAKE A CONSTRCOOR
//      chage     Transitions  map[string]map[sting]string to     Transitions  map[string]map[rune]string
//      make te dfa struct 
//      mk
