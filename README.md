# MCTS
An efficient Go implemetation of the Monte Carlo Tree Search algorithm

More than 20 times faster than the [Python implementation](https://github.com/pbsinclair42/MCTS) by Paul Sinclair et al. 

## How to Use

1. Provide a custom environment

Implement all methods of the `State` interface. Specifically:

```go
type State interface {
  GetCurrentPlayer() int
  GetPossibleActions() []any
  TakeAction(action any) State
  IsTerminal() bool
  GetReward() int
}
```

where `action` can be any hashable type. We provide example environments under `env/`.

2. Search for the best action given the current state

An example using the toy environment:
```go
import (
  "github.com/Zacchaeus14/MCTS"
  "github.com/Zacchaeus14/MCTS/policy"
)
```

## TODO
- [X] Modularization
- [ ] Parallel rollout
