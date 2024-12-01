package nfa

import (
	"testing"
)

func TestValidateStringDac(t *testing.T) {
	// Define the first NFA
	nfa1 := &NFA{
		StartState: &StateNode{
			StateName:   "q0",
			IsAccepting: false,
			Transitions: map[rune][]*StateNode{
				'a': {&StateNode{StateName: "q1", IsAccepting: true}},
			},
		},
	}

	// Define the second NFA
	nfa2 := &NFA{
		StartState: &StateNode{
			StateName:   "s0",
			IsAccepting: false,
			Transitions: map[rune][]*StateNode{
				'b': {&StateNode{StateName: "s1", IsAccepting: false, Transitions: map[rune][]*StateNode{
					'c': {&StateNode{StateName: "s2", IsAccepting: true}},
				}}},
			},
		},
	}

	testCases := []struct {
		nfa      *NFA
		input    []rune
		expected bool
	}{
		// Test cases for the first NFA
		{nfa: nfa1, input: []rune("a"), expected: true},
		{nfa: nfa1, input: []rune("aa"), expected: false},
		{nfa: nfa1, input: []rune(""), expected: false},
		{nfa: nfa1, input: []rune("b"), expected: false},
		{nfa: nfa1, input: []rune("ab"), expected: false},
		{nfa: nfa1, input: []rune("a"), expected: true},
		{nfa: nfa1, input: []rune("a"), expected: true},
		{nfa: nfa1, input: []rune("a"), expected: true},
		{nfa: nfa1, input: []rune("a"), expected: true},
		{nfa: nfa1, input: []rune("a"), expected: true},

		// Test cases for the second NFA
		{nfa: nfa2, input: []rune("bc"), expected: true},
		{nfa: nfa2, input: []rune("b"), expected: false},
		{nfa: nfa2, input: []rune(""), expected: false},
		{nfa: nfa2, input: []rune("c"), expected: false},
		{nfa: nfa2, input: []rune("abc"), expected: false},
		{nfa: nfa2, input: []rune("bc"), expected: true},
		{nfa: nfa2, input: []rune("bc"), expected: true},
		{nfa: nfa2, input: []rune("bc"), expected: true},
		{nfa: nfa2, input: []rune("bc"), expected: true},
		{nfa: nfa2, input: []rune("bc"), expected: true},
	}

	for _, tc := range testCases {
		result := tc.nfa.ValidateString(string(tc.input))
		if result != tc.expected {
			t.Errorf("ValidateStringDac(%q) = %v; want %v", string(tc.input), result, tc.expected)
		}
	}
}

//how to run the test cases in the terminal as test file is in the nfa package
//go test -v
//go test -v -run TestValidateStringDac
//go test -v -run TestValidateStringDac -count=1
//go test -v -run TestValidateStringDac -count=1 -cover
//go test -v -run TestValidateStringDac -count=1 -coverprofile=coverage.out
//go tool cover -html=coverage.out
