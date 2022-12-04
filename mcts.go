package MCTS

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type State interface {
	GetCurrentPlayer() int
	GetPossibleActions() []any
	TakeAction(a any) State
	IsTerminal() bool
	GetReward() int
}

type TreeNode struct {
	State           State
	IsTerminal      bool
	IsFullyExpanded bool
	Parent          *TreeNode
	NumVisits       int
	TotalReward     int
	Children        map[any]*TreeNode
}

func NewTreeNode(state State, parent *TreeNode) *TreeNode {
	node := TreeNode{}
	node.State = state
	node.IsTerminal = state.IsTerminal()
	node.IsFullyExpanded = node.IsTerminal
	node.Parent = parent
	node.Children = map[any]*TreeNode{}
	return &node
}

type MCTS struct {
	TimeLimit           int
	IterationLimit      int
	LimitType           string
	ExplorationConstant float64
	Rollout             func(State, chan int)
	Root                *TreeNode
	Jobs                int
}

func NewMCTS(timeLimit int, iterationLimit int, explorationConstant float64, rollout func(State, chan int), Jobs int) *MCTS {
	mcts := MCTS{}
	if timeLimit > 0 {
		if iterationLimit > 0 {
			panic("Cannot have both a time limit and an iteration limit")
		}
		mcts.TimeLimit = timeLimit
		mcts.LimitType = "time"
	} else {
		if iterationLimit == 0 {
			panic("Must have either a time limit or an iteration limit")
		}
		mcts.IterationLimit = iterationLimit
		mcts.LimitType = "iterations"
	}
	mcts.ExplorationConstant = explorationConstant
	mcts.Rollout = rollout
	mcts.Jobs = Jobs
	return &mcts
}

func (mcts *MCTS) Search(initialState State, verbose int) any {
	mcts.Root = NewTreeNode(initialState, nil)
	if mcts.LimitType == "time" {
		timeLimit := time.Now().UnixNano()/1000000 + int64(mcts.TimeLimit)
		for time.Now().UnixNano()/1000000 < timeLimit {
			mcts.executeRound()
		}
	} else {
		for i := 0; i < mcts.IterationLimit; i++ {
			mcts.executeRound()
		}
	}
	bestChild := mcts.getBestChild(mcts.Root, 0)
	bestMeanReward := getMeanReward(bestChild)
	if verbose == 2 {
		for action, child := range mcts.Root.Children {
			fmt.Printf("Action: %v\t%.2f%% Visits\t%.3f%% Wins\n", action, 100*float64(child.NumVisits)/float64(mcts.Root.NumVisits), 100*(float64(child.TotalReward)/float64(child.NumVisits)+1)/2)
		}
	}
	for action, child := range mcts.Root.Children {
		if getMeanReward(child) == bestMeanReward {
			if verbose == 1 {
				fmt.Printf("Action: %v\t%.2f%% Visits\t%.3f%% Wins\n", action, 100*float64(child.NumVisits)/float64(mcts.Root.NumVisits), 100*(float64(child.TotalReward)/float64(child.NumVisits)+1)/2)
			}
			return action
		}
	}
	panic("Should never reach here")
}

func (mcts *MCTS) executeRound() {
	node := mcts.selectNode(mcts.Root)
	reward := 0
	ch := make(chan int)
	for i := 0; i < mcts.Jobs; i++ {
		go mcts.Rollout(node.State, ch)
	}
	for i := 0; i < mcts.Jobs; i++ {
		reward += <-ch
	}
	mcts.backpropogate(node, reward)
}

func (mcts *MCTS) selectNode(node *TreeNode) *TreeNode {
	for !node.IsTerminal {
		if node.IsFullyExpanded {
			node = mcts.getBestChild(node, mcts.ExplorationConstant)
		} else {
			return mcts.expand(node)
		}
	}
	return node
}

func (mcts *MCTS) expand(node *TreeNode) *TreeNode {
	actions := node.State.GetPossibleActions()
	for _, action := range actions {
		if _, ok := node.Children[action]; !ok {
			newNode := NewTreeNode(node.State.TakeAction(action), node)
			node.Children[action] = newNode
			if len(actions) == len(node.Children) {
				node.IsFullyExpanded = true
			}
			return newNode
		}
	}
	panic("Should never reach here")
}

func (mcts *MCTS) backpropogate(node *TreeNode, reward int) {
	for node != nil {
		node.NumVisits += mcts.Jobs
		node.TotalReward += reward
		node = node.Parent
	}
}

func (mcts *MCTS) getBestChild(node *TreeNode, explorationValue float64) *TreeNode {
	bestValue := -math.MaxFloat64
	bestNodes := []*TreeNode{}
	for _, child := range node.Children {
		nodeValue := float64(node.State.GetCurrentPlayer()*child.TotalReward)/float64(child.NumVisits) + explorationValue*math.Sqrt(2.0*math.Log(float64(node.NumVisits))/float64(child.NumVisits))
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
	return float64(node.TotalReward) / float64(node.NumVisits)
}
