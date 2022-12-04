package env

import (
	"fmt"
	"github.com/Zacchaeus14/MCTS"
	"github.com/Zacchaeus14/MCTS/policy"
	"math"
	"testing"
)

func TestNewState(t *testing.T) {
	NewNaughtsAndCrossesState()
}

func TestTerminal(t *testing.T) {
	tState := NaughtsAndCrossesState{[][]int{{1, 1, 1}, {0, 0, 0}, {0, 0, 0}}, 1}
	if !tState.IsTerminal() {
		t.Fatalf("%v should be terminal state", tState)
	}
	tState1 := NaughtsAndCrossesState{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 1}
	if !tState1.IsTerminal() {
		t.Fatalf("%v should be terminal state", tState1)
	}
	tState2 := NaughtsAndCrossesState{[][]int{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}}, 1}
	if !tState2.IsTerminal() {
		t.Fatalf("%v should be terminal state", tState2)
	}
	tState3 := NaughtsAndCrossesState{[][]int{{1, 0, 0}, {1, 0, 0}, {1, 0, 0}}, 1}
	if !tState3.IsTerminal() {
		t.Fatalf("%v should be terminal state", tState3)
	}
}

func TestMCTS(t *testing.T) {
	initialState := NewNaughtsAndCrossesState()
	searcher := MCTS.NewMCTS(1000, 0, 1/math.Sqrt(2), policy.ParallelRandomPolicy, 10)
	bestAction := searcher.Search(initialState, true).(NaughtsAndCrossesAction)
	targetAction := NaughtsAndCrossesAction{1, 1, 1}
	if bestAction != targetAction {
		t.Fatalf("Best action should be %v, but has %v", targetAction, bestAction)
	}
	fmt.Println("Total rounds searched in 1 second:", searcher.Root.NumVisits)
	middleState := &NaughtsAndCrossesState{[][]int{{0, 0, 0}, {0, 1, 0}, {0, -1, 0}}, 1}
	searcher = MCTS.NewMCTS(1000, 0, 10, policy.ParallelRandomPolicy, 1)
	bestAction = searcher.Search(middleState, true).(NaughtsAndCrossesAction)
}
