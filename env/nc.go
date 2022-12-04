package env

import "github.com/Zacchaeus14/MCTS"

type NaughtsAndCrossesAction struct {
	Player int
	X      int
	Y      int
}

type NaughtsAndCrossesState struct {
	Board         [][]int
	CurrentPlayer int
}

func NewNaughtsAndCrossesState() *NaughtsAndCrossesState {
	return &NaughtsAndCrossesState{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 1}
}

func (s *NaughtsAndCrossesState) GetCurrentPlayer() int {
	return s.CurrentPlayer
}

func (s *NaughtsAndCrossesState) GetPossibleActions() []any {
	possibleActions := []any{}
	for i := range s.Board {
		for j := range s.Board[i] {
			if s.Board[i][j] == 0 {
				possibleActions = append(possibleActions, NaughtsAndCrossesAction{s.CurrentPlayer, i, j})
			}
		}
	}
	return possibleActions
}

func (s *NaughtsAndCrossesState) TakeAction(a any) MCTS.State {
	newState := NewNaughtsAndCrossesState()
	for i, row := range s.Board {
		copy(newState.Board[i], row)
	}
	ncAction := a.(NaughtsAndCrossesAction)
	newState.Board[ncAction.X][ncAction.Y] = ncAction.Player
	newState.CurrentPlayer = -s.CurrentPlayer
	return newState
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
func (s *NaughtsAndCrossesState) IsTerminal() bool {
	sum := 0
	for _, row := range s.Board {
		sum = 0
		for _, e := range row {
			sum += e
		}
		if abs(sum) == 3 {
			return true
		}
	}
	for j := range s.Board[0] {
		sum = 0
		for i := range s.Board {
			sum += s.Board[i][j]
		}
		if abs(sum) == 3 {
			return true
		}
	}
	sum = 0
	for i := range s.Board {
		sum += s.Board[i][i]
	}
	if abs(sum) == 3 {
		return true
	}
	sum = 0
	for i := range s.Board {
		sum += s.Board[i][len(s.Board)-i-1]
	}
	if abs(sum) == 3 {
		return true
	}
	for _, row := range s.Board {
		for _, e := range row {
			if e == 0 {
				return false
			}
		}
	}
	return true
}

func (s *NaughtsAndCrossesState) GetReward() int {
	sum := 0
	for _, row := range s.Board {
		sum = 0
		for _, e := range row {
			sum += e
		}
		if abs(sum) == 3 {
			return sum / 3
		}
	}
	for j := range s.Board[0] {
		sum = 0
		for i := range s.Board {
			sum += s.Board[i][j]
		}
		if abs(sum) == 3 {
			return sum / 3
		}
	}
	sum = 0
	for i := range s.Board {
		sum += s.Board[i][i]
	}
	if abs(sum) == 3 {
		return sum / 3
	}
	sum = 0
	for i := range s.Board {
		sum += s.Board[i][len(s.Board)-i-1]
	}
	if abs(sum) == 3 {
		return sum / 3
	}
	return 0
}
