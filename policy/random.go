package policy

import (
	"github.com/Zacchaeus14/MCTS"
	"math/rand"
)

func RandomPolicy(state MCTS.State) int {
	for !state.IsTerminal() {
		actions := state.GetPossibleActions()
		action := actions[rand.Intn(len(actions))]
		state = state.TakeAction(action)
	}
	return state.GetReward()
}
