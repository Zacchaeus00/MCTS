package env

import (
	"fmt"
	"github.com/Zacchaeus14/MCTS"
	"github.com/Zacchaeus14/MCTS/policy"
	"testing"
)

func TestNewC4State(t *testing.T) {
	state := NewConnect4StateState()
	fmt.Println(state)
}

func TestC4MCTS(t *testing.T) {
	initialState := NewConnect4StateState()
	searcher := MCTS.NewMCTS(1000, 0, 10, policy.ParallelRandomPolicy, 10)
	bestAction := searcher.Search(initialState, 2).(Connect4Action)
	targetAction := Connect4Action{1, 3}
	if bestAction != targetAction {
		t.Fatalf("Best action should be %v, but has %v", targetAction, bestAction)
	}
	fmt.Println("Total rounds searched in 1 second:", searcher.Root.NumVisits)
}
