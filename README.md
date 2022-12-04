# MCTS
An efficient Go implemetation of the Monte Carlo Tree Search algorithm with a friendly interface

More than 20/60 times faster single/multi thread than the [Python implementation](https://github.com/pbsinclair42/MCTS) by Paul Sinclair et al. 

## How to Use

1. Provide a custom environment. Implement all methods of the `State` interface below, where `action` can be any hashable type. We provide example environments under `env/`.

```go
type State interface {
  GetCurrentPlayer() int
  GetPossibleActions() []any
  TakeAction(action any) State
  IsTerminal() bool
  GetReward() int
}
```

2. Search for the best action given the current state. You can define a custom rollout policy or use the default random one. Below is an example with the toy environment.
```go
import (
  "github.com/Zacchaeus14/MCTS"
  "github.com/Zacchaeus14/MCTS/policy"
  "math"
)
initialState := NewNaughtsAndCrossesState()
searcher := MCTS.NewMCTS(1000, 0, 1.0/math.Sqrt(2), policy.RandomPolicy) // limit search time to one second
bestAction := searcher.Search(initialState) // {1, 1, 1}
```

## TODO
- [X] Modularization
- [X] Parallel rollout
