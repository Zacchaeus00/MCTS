package env

import "github.com/Zacchaeus14/MCTS"

type Connect4Action struct {
	player int
	x      int
}

type Connect4State struct {
	board         [][]int
	currentPlayer int
}

func NewConnect4StateState() *Connect4State {
	state := Connect4State{}
	state.board = make([][]int, 6)
	for i := 0; i < 6; i++ {
		state.board[i] = make([]int, 7)
	}
	state.currentPlayer = 1
	return &state
}

func (s *Connect4State) GetCurrentPlayer() int {
	return s.currentPlayer
}

func (s *Connect4State) GetPossibleActions() []any {
	possibleActions := []any{}
	for j, e := range s.board[0] {
		if e == 0 {
			possibleActions = append(possibleActions, Connect4Action{s.currentPlayer, j})
		}
	}
	return possibleActions
}

func (s *Connect4State) TakeAction(a any) MCTS.State {
	newState := NewConnect4StateState()
	for i, row := range s.board {
		copy(newState.board[i], row)
	}
	c4Action := a.(Connect4Action)
	for i := 5; i >= 0; i-- {
		if newState.board[i][c4Action.x] == 0 {
			newState.board[i][c4Action.x] = c4Action.player
			break
		}
	}
	newState.currentPlayer = -s.currentPlayer
	return newState
}

func (s *Connect4State) IsTerminal() bool {
	return s.checkFull() || s.checkWin(1) || s.checkWin(-1)
}

func (s *Connect4State) GetReward() int {
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

func (s *Connect4State) checkFull() bool {
	for _, row := range s.board {
		for _, e := range row {
			if e == 0 {
				return false
			}
		}
	}
	return true
}

func (s *Connect4State) checkWin(player int) bool {
	h, w := len(s.board), len(s.board[0])
	// horizontalCheck
	for j := 0; j < w-3; j++ {
		for i := 0; i < h; i++ {
			if s.board[i][j] == player && s.board[i][j+1] == player && s.board[i][j+2] == player && s.board[i][j+3] == player {
				return true
			}
		}
	}
	// verticalCheck
	for i := 0; i < h-3; i++ {
		for j := 0; j < w; j++ {
			if s.board[i][j] == player && s.board[i+1][j] == player && s.board[i+2][j] == player && s.board[i+3][j] == player {
				return true
			}
		}
	}
	// ascendingDiagonalCheck
	for i := 3; i < h; i++ {
		for j := 0; j < w-3; j++ {
			if s.board[i][j] == player && s.board[i-1][j+1] == player && s.board[i-2][j+2] == player && s.board[i-3][j+3] == player {
				return true
			}
		}
	}
	// descendingDiagonalCheck
	for i := 3; i < h; i++ {
		for j := 3; j < w; j++ {
			if s.board[i][j] == player && s.board[i-1][j-1] == player && s.board[i-2][j-2] == player && s.board[i-3][j-3] == player {
				return true
			}
		}
	}
	return false
}
