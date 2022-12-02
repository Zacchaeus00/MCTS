package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Action struct {
	player int
	x      int
	y      int
}

type State interface {
	getCurrentPlayer() int
	getPossibleActions() []Action
	takeAction(a Action) State
	isTerminal() bool
	getReward() int
}

type NaughtsAndCrossesState struct {
	board         [][]int
	currentPlayer int
}

func initNaughtsAndCrossesState() *NaughtsAndCrossesState {
	return &NaughtsAndCrossesState{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 1}
}

func (s *NaughtsAndCrossesState) getCurrentPlayer() int {
	return s.currentPlayer
}

func (s *NaughtsAndCrossesState) getPossibleActions() []Action {
	possibleActions := []Action{}
	for i := range s.board {
		for j := range s.board[i] {
			if s.board[i][j] == 0 {
				possibleActions = append(possibleActions, Action{s.currentPlayer, i, j})
			}
		}
	}
	return possibleActions
}

func (s *NaughtsAndCrossesState) takeAction(a Action) *NaughtsAndCrossesState {
	newState := *initNaughtsAndCrossesState()
	for i, row := range s.board {
		copy(newState.board[i], row)
	}
	newState.board[a.x][a.y] = a.player
	newState.currentPlayer = -s.currentPlayer
	return &newState
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func (s *NaughtsAndCrossesState) isTerminal() bool {
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

func (s *NaughtsAndCrossesState) getReward() int {
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

// func main() {
// 	state := initNaughtsAndCrossesState()
// 	fmt.Println(state)

// 	action := Action{1, 0, 0}
// 	action1 := Action{1, 0, 0}
// 	newState := state.takeAction(action)
// 	fmt.Println(state)
// 	fmt.Println(newState)
// 	set := map[Action]int{}
// 	set[action] = 1
// 	fmt.Println(action1 == action)
// 	fmt.Println(state.isTerminal())
// 	tState := NaughtsAndCrossesState{[][]int{{1, 1, 1}, {0, 0, 0}, {0, 0, 0}}, 1}
// 	fmt.Println(tState.isTerminal())
// 	tState1 := NaughtsAndCrossesState{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 1}
// 	fmt.Println(tState1.isTerminal())
// 	tState2 := NaughtsAndCrossesState{[][]int{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}}, 1}
// 	fmt.Println(tState2.isTerminal())
// 	tState3 := NaughtsAndCrossesState{[][]int{{1, 0, 0}, {1, 0, 0}, {1, 0, 0}}, 1}
// 	fmt.Println(tState3.isTerminal())
// }

func randomPolicy(state *NaughtsAndCrossesState) int {
	for !state.isTerminal() {
		actions := state.getPossibleActions()
		action := actions[rand.Intn(len(actions))]
		state = state.takeAction(action)
	}
	return state.getReward()
}

type TreeNode struct {
	state           NaughtsAndCrossesState
	isTerminal      bool
	isFullyExpanded bool
	parent          *TreeNode
	numVisits       int
	totalReward     int
	children        map[Action]*TreeNode
}

func initTreeNode(state *NaughtsAndCrossesState, parent *TreeNode) *TreeNode {
	node := TreeNode{}
	node.state = *state
	node.isTerminal = state.isTerminal()
	node.isFullyExpanded = node.isTerminal
	node.parent = parent
	node.children = map[Action]*TreeNode{}
	return &node
}

type rollout_t func(*NaughtsAndCrossesState) int

type MCTS struct {
	timeLimit           int
	iterationLimit      int
	limitType           string
	explorationConstant float64
	rollout             rollout_t
	root                *TreeNode
}

func initMCTS(timeLimit int, iterationLimit int, explorationConstant float64, rollout rollout_t) *MCTS {
	mcts := MCTS{}
	if timeLimit > 0 {
		if iterationLimit > 0 {
			panic("Cannot have both a time limit and an iteration limit")
		}
		mcts.timeLimit = timeLimit
		mcts.limitType = "time"
	} else {
		if iterationLimit == 0 {
			panic("Must have either a time limit or an iteration limit")
		}
		mcts.iterationLimit = iterationLimit
		mcts.limitType = "iterations"
	}
	mcts.explorationConstant = explorationConstant
	mcts.rollout = rollout
	return &mcts
}

func (self *MCTS) search(initialState *NaughtsAndCrossesState) Action {
	self.root = initTreeNode(initialState, nil)
	if self.limitType == "time" {
		timeLimit := time.Now().UnixNano()/1000000 + int64(self.timeLimit)
		for time.Now().UnixNano()/1000000 < timeLimit {
			self.executeRound()
		}
	} else {
		for i := 0; i < self.iterationLimit; i++ {
			self.executeRound()
		}
	}
	bestChild := self.getBestChild(self.root, 0)
	bestMeanReward := getMeanReward(bestChild)
	for action, child := range self.root.children {
		if getMeanReward(child) == bestMeanReward {
			return action
		}
	}
	panic("Should never reach here")
}

func (self *MCTS) executeRound() {
	node := self.selectNode(self.root)
	reward := self.rollout(&node.state)
	self.backpropogate(node, reward)
}

func (self *MCTS) selectNode(node *TreeNode) *TreeNode {
	for !node.isTerminal {
		if node.isFullyExpanded {
			node = self.getBestChild(node, self.explorationConstant)
		} else {
			return self.expand(node)
		}
	}
	return node
}

func (self *MCTS) expand(node *TreeNode) *TreeNode {
	actions := node.state.getPossibleActions()
	for _, action := range actions {
		if _, ok := node.children[action]; !ok {
			newNode := initTreeNode(node.state.takeAction(action), node)
			node.children[action] = newNode
			if len(actions) == len(node.children) {
				node.isFullyExpanded = true
			}
			return newNode
		}
	}
	panic("Should never reach here")
}

func (self *MCTS) backpropogate(node *TreeNode, reward int) {
	for node != nil {
		node.numVisits++
		node.totalReward += reward
		node = node.parent
	}
}

func (self *MCTS) getBestChild(node *TreeNode, explorationValue float64) *TreeNode {
	bestValue := -math.MaxFloat64
	bestNodes := []*TreeNode{}
	for _, child := range node.children {
		nodeValue := float64(node.state.getCurrentPlayer()*child.totalReward)/float64(child.numVisits) + explorationValue*math.Sqrt(2.0*math.Log(float64(node.numVisits))/float64(child.numVisits))
		if nodeValue > bestValue {
			bestValue = nodeValue
			bestNodes = []*TreeNode{child}
		} else if nodeValue == bestValue {
			bestNodes = append(bestNodes, child)
		}
	}
	return bestNodes[rand.Intn(len(bestNodes))]
}

func getMeanReward(node *TreeNode) float64 {
	return float64(node.totalReward) / float64(node.numVisits)
}

func main() {
	initialState := initNaughtsAndCrossesState()
	searcher := initMCTS(1000, 0, 1.0/math.Sqrt(2), randomPolicy)
	bestAction := searcher.search(initialState)
	fmt.Println(bestAction)
	fmt.Println(searcher.root.numVisits)
}
