package env

import (
	"fmt"
	"github.com/Zacchaeus14/MCTS"
	"math/rand"
	"testing"
	"time"
)

func TestNewGomokuState(t *testing.T) {
	state := NewGomokuState()
	fmt.Println(state)
}

func TestNewGomokuPossibleActions(t *testing.T) {
	var state MCTS.State = NewGomokuState()
	actions := state.GetPossibleActions()
	if len(actions) != 1 {
		t.Fatalf("State %v get unmatched number of actions %v", state, len(actions))
	}
	state = state.TakeAction(GomokuAction{1, 7, 7})
	actions = state.GetPossibleActions()
	if len(actions) != 5*5-1 {
		t.Fatalf("State %v get unmatched number of actions %v", state, len(actions))
	}
	state = state.TakeAction(GomokuAction{1, 7, 8})
	actions = state.GetPossibleActions()
	if len(actions) != 6*5-2 {
		t.Fatalf("State %v get unmatched number of actions %v", state, len(actions))
	}
	fmt.Println(state.(*GomokuState).minX, state.(*GomokuState).minY, state.(*GomokuState).maxX, state.(*GomokuState).maxY)
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, e := range row {
			msg := "_"
			if e == 1 {
				msg = "x"
			}
			if e == -1 {
				msg = "o"
			}
			print(msg, "\t")
		}
		println()
	}
	for i := range board[0] {
		print(i, "\t")
	}
	println()
}

func TestGomokuRandomPlay(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var state MCTS.State = NewGomokuState()
	for !state.IsTerminal() {
		actions := state.GetPossibleActions()
		state = state.TakeAction(actions[rand.Intn(len(actions))])
	}
	printBoard(state.(*GomokuState).Board)
	fmt.Println(state.GetReward())
}
