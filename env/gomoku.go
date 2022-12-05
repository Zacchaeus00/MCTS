package env

import (
	"github.com/Zacchaeus14/MCTS"
)

type GomokuAction struct {
	Player int
	X      int
	Y      int
}

type GomokuState struct {
	Board         [][]int
	CurrentPlayer int
	minX          int
	minY          int
	maxX          int
	maxY          int
	empty         bool
}

func NewGomokuState() *GomokuState {
	state := GomokuState{}
	state.Board = make([][]int, 15)
	for i := 0; i < 15; i++ {
		state.Board[i] = make([]int, 15)
	}
	state.CurrentPlayer = 1
	state.minX = 14
	state.minY = 14
	state.maxX = 0
	state.maxY = 0
	state.empty = true
	return &state
}

func (s *GomokuState) GetCurrentPlayer() int {
	return s.CurrentPlayer
}

func (s *GomokuState) isValid(i int, j int) bool {
	if i < 0 || i >= len(s.Board) || j < 0 || j >= len(s.Board[0]) {
		return false
	}
	if s.Board[i][j] == 0 {
		return true
	}
	return false
}

func (s *GomokuState) GetPossibleActions() []any {
	if s.empty {
		return []any{GomokuAction{1, 7, 7}}
	}
	possibleActions := []any{}
	for i := s.minX - 2; i <= s.maxX+2; i++ {
		for j := s.minY - 2; j <= s.maxY+2; j++ {
			if s.isValid(i, j) {
				possibleActions = append(possibleActions, GomokuAction{s.CurrentPlayer, i, j})
			}
		}
	}
	return possibleActions
}

func (s *GomokuState) TakeAction(a any) MCTS.State {
	newState := NewGomokuState()
	for i, row := range s.Board {
		copy(newState.Board[i], row)
	}
	gomokuAction := a.(GomokuAction)
	newState.Board[gomokuAction.X][gomokuAction.Y] = gomokuAction.Player
	newState.CurrentPlayer = -s.CurrentPlayer
	newState.minX = s.minX
	newState.maxX = s.maxX
	newState.minY = s.minY
	newState.maxY = s.maxY
	if a.(GomokuAction).X < newState.minX {
		newState.minX = a.(GomokuAction).X
	}
	if a.(GomokuAction).Y < newState.minY {
		newState.minY = a.(GomokuAction).Y
	}
	if a.(GomokuAction).X > newState.maxX {
		newState.maxX = a.(GomokuAction).X
	}
	if a.(GomokuAction).Y > newState.maxY {
		newState.maxY = a.(GomokuAction).Y
	}
	newState.empty = false
	return newState
}

func (s *GomokuState) IsTerminal() bool {
	return s.checkFull() || s.checkWin(1) || s.checkWin(-1)
}

func (s *GomokuState) GetReward() int {
	if s.checkFull() {
		return 0
	}
	if s.checkWin(1) {
		return 1
	}
	if s.checkWin(-1) {
		return -1
	}
	panic("Shouldn't reach here")
}

func (s *GomokuState) checkFull() bool {
	for _, row := range s.Board {
		for _, e := range row {
			if e == 0 {
				return false
			}
		}
	}
	return true
}

func (s *GomokuState) checkWin(player int) bool {
	h, w := len(s.Board), len(s.Board[0])
	// horizontalCheck
	for j := 0; j < w-4; j++ {
		for i := 0; i < h; i++ {
			if s.Board[i][j] == player && s.Board[i][j+1] == player && s.Board[i][j+2] == player && s.Board[i][j+3] == player && s.Board[i][j+4] == player {
				return true
			}
		}
	}
	// verticalCheck
	for i := 0; i < h-4; i++ {
		for j := 0; j < w; j++ {
			if s.Board[i][j] == player && s.Board[i+1][j] == player && s.Board[i+2][j] == player && s.Board[i+3][j] == player && s.Board[i+4][j] == player {
				return true
			}
		}
	}
	// ascendingDiagonalCheck
	for i := 4; i < h; i++ {
		for j := 0; j < w-4; j++ {
			if s.Board[i][j] == player && s.Board[i-1][j+1] == player && s.Board[i-2][j+2] == player && s.Board[i-3][j+3] == player && s.Board[i-3][j+4] == player {
				return true
			}
		}
	}
	// descendingDiagonalCheck
	for i := 4; i < h; i++ {
		for j := 4; j < w; j++ {
			if s.Board[i][j] == player && s.Board[i-1][j-1] == player && s.Board[i-2][j-2] == player && s.Board[i-3][j-3] == player && s.Board[i-4][j-4] == player {
				return true
			}
		}
	}
	return false
}
