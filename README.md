# MCTS
An efficient Go implemetation of the Monte Carlo Tree Search algorithm

More than 20 times faster than the [Python implementation](https://github.com/pbsinclair42/MCTS) by Paul Sinclair et al. 

## How to Use

Implement all methods of the `State` interface. Specifically,

```go
type State interface {
  GetCurrentPlayer() int
  GetPossibleActions() []any
  TakeAction(a any) State
  IsTerminal() bool
  GetReward() int
}
```

## TODO
- [X] Modularization
- [ ] Parallel rollout
