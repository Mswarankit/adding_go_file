package elevator

import "sync"

type Direction int

const (
	Idle Direction = iota
	Up
	Down
)

type Request struct {
	CurrentFloor int
	DesiredFloor int
	IsInternal   bool
}

type State struct {
	CurrentFloor int
	Direction    Direction
	IsMoving     bool
	DoorsOpen    bool
}

type Elevator struct {
	state     State
	requests  *RequestQueue
	mutex     sync.Mutex
	maxFloor  int
	minFloor  int
	emergency bool
}

type RequestQueue struct {
	upQueue   []Request
	downQueue []Request
	mutex     sync.Mutex
}
