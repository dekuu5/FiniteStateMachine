package dfa



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
