package env

import "github.com/Zacchaeus14/MCTS"

type NaughtsAndCrossesAction struct {
	player int
	x      int
	y      int
}

type NaughtsAndCrossesState struct {
	board         [][]int
	currentPlayer int
}

func NewNaughtsAndCrossesState() *NaughtsAndCrossesState {
	return &NaughtsAndCrossesState{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 1}
}

func (s *NaughtsAndCrossesState) GetCurrentPlayer() int {
	return s.currentPlayer
}

func (s *NaughtsAndCrossesState) GetPossibleActions() []any {
	possibleActions := []any{}
	for i := range s.board {
		for j := range s.board[i] {
			if s.board[i][j] == 0 {
				possibleActions = append(possibleActions, NaughtsAndCrossesAction{s.currentPlayer, i, j})
			}
		}
	}
	return possibleActions
}

func (s *NaughtsAndCrossesState) TakeAction(a any) MCTS.State {
	newState := NewNaughtsAndCrossesState()
	for i, row := range s.board {
		copy(newState.board[i], row)
	}
	ncAction := a.(NaughtsAndCrossesAction)
	newState.board[ncAction.x][ncAction.y] = ncAction.player
	newState.currentPlayer = -s.currentPlayer
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
	for _, row := range s.board {
		sum = 0
		for _, e := range row {
			sum += e
		}
		if abs(sum) == 3 {
			return true
		}
	}
	for j := range s.board[0] {
		sum = 0
		for i := range s.board {
			sum += s.board[i][j]
		}
		if abs(sum) == 3 {
			return true
		}
	}
	sum = 0
	for i := range s.board {
		sum += s.board[i][i]
	}
	if abs(sum) == 3 {
		return true
	}
	sum = 0
	for i := range s.board {
		sum += s.board[i][len(s.board)-i-1]
	}
	if abs(sum) == 3 {
		return true
	}
	for _, row := range s.board {
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
	for _, row := range s.board {
		sum = 0
		for _, e := range row {
			sum += e
		}
		if abs(sum) == 3 {
			return sum / 3
		}
	}
	for j := range s.board[0] {
		sum = 0
		for i := range s.board {
			sum += s.board[i][j]
		}
		if abs(sum) == 3 {
			return sum / 3
		}
	}
	sum = 0
	for i := range s.board {
		sum += s.board[i][i]
	}
	if abs(sum) == 3 {
		return sum / 3
	}
	sum = 0
	for i := range s.board {
		sum += s.board[i][len(s.board)-i-1]
	}
	if abs(sum) == 3 {
		return sum / 3
	}
	return 0
}
