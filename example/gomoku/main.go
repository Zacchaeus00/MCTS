package main

import (
	"fmt"
	"github.com/Zacchaeus14/MCTS"
	"github.com/Zacchaeus14/MCTS/env"
	"github.com/Zacchaeus14/MCTS/policy"
)

func printBoard(board [][]int) {
	for i, row := range board {
		print(i, "\t")
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
	print("\t")
	for i := range board[0] {
		print(i, "\t")
	}
	println()
}

func pve() {
	var state MCTS.State
	state = env.NewGomokuState()

	var first string
	fmt.Print("go first (y/n): ")
	fmt.Scan(&first)
	player := 1
	if first == "n" {
		searcher := MCTS.NewMCTS(1000, 0, 10, policy.ParallelRandomPolicy, 10)
		bestAction := searcher.Search(state, 1)
		state = state.TakeAction(bestAction)
		player = -1
	}

	for !state.IsTerminal() {
		printBoard(state.(*env.GomokuState).Board)
		x := 0
		y := 0
		fmt.Print("input x: ")
		fmt.Scan(&x)
		fmt.Print("input y: ")
		fmt.Scan(&y)
		state = state.TakeAction(env.GomokuAction{Player: player, X: x, Y: y})
		if state.IsTerminal() {
			break
		}
		searcher := MCTS.NewMCTS(3000, 0, 1, policy.ParallelRandomPolicy, 10)
		bestAction := searcher.Search(state, 1)
		state = state.TakeAction(bestAction)
	}
	printBoard(state.(*env.GomokuState).Board)
}

func selfPlay(n int) {
	total_reward := 0
	for i := 0; i < n; i++ {
		var state MCTS.State
		state = env.NewConnect4State()

		for !state.IsTerminal() {
			searcher := MCTS.NewMCTS(1000, 0, 5, policy.ParallelRandomPolicy, 10)
			bestAction := searcher.Search(state, 0)
			state = state.TakeAction(bestAction)
		}
		reward := state.GetReward()
		total_reward += reward
		fmt.Println(i, reward, total_reward)
	}
}

func main() {
	//selfPlay(100)
	pve()
}
