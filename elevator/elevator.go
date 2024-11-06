package elevator

import (
	"fmt"
	"time"
)

func NewElevator(minFloor, maxFloor int) *Elevator {
	return &Elevator{
		state: State{
			CurrentFloor: 1,
			Direction:    Idle,
			IsMoving:     false,
			DoorsOpen:    false,
		},
		requests: &RequestQueue{
			upQueue:   make([]Request, 0),
			downQueue: make([]Request, 0),
		},
		maxFloor:  maxFloor,
		minFloor:  minFloor,
		emergency: false,
	}
}

func (e *Elevator) AddRequest(req Request) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if req.DesiredFloor > e.maxFloor || req.DesiredFloor < e.minFloor {
		return fmt.Errorf("invalid floor request: %d", req.DesiredFloor)
	}

	if req.DesiredFloor > e.state.CurrentFloor {
		e.requests.addUpRequest(req)
	} else if req.DesiredFloor < e.state.CurrentFloor {
		e.requests.addDownRequest(req)
	}
	return nil
}

func (e *Elevator) Start() {
	go func() {
		for {
			if e.emergency {
				e.handleEmergency()
				continue
			}
			e.processRequest()
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func (e *Elevator) processRequest() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if len(e.requests.upQueue) > 0 && (e.state.Direction == Up || e.state.Direction == Idle) {
		e.processUpQueue()
	} else if len(e.requests.downQueue) > 0 && (e.state.Direction == Down || e.state.Direction == Idle) {
		e.processDownQueue()
	} else {
		e.state.Direction = Idle
		e.state.IsMoving = false
	}
}

func (e *Elevator) GetStatus() State {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	return e.state
}

func (e *Elevator) TriggerEmergency() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.emergency = true
	e.state.DoorsOpen = true
	e.state.IsMoving = false
	e.state.Direction = Idle
}

func (e *Elevator) processUpQueue() {
	if len(e.requests.upQueue) == 0 {
		return
	}
	e.state.Direction = Up
	e.state.IsMoving = true
	nextRequest := e.requests.upQueue[0]

	if e.state.CurrentFloor == nextRequest.DesiredFloor {
		e.requests.upQueue = e.requests.upQueue[1:]

		e.state.IsMoving = false
		e.state.DoorsOpen = true
		time.Sleep(10 * time.Second)
		e.state.DoorsOpen = false
	} else {
		e.state.CurrentFloor++
		time.Sleep(1 * time.Second)
	}

	if e.state.CurrentFloor == e.maxFloor {
		e.state.Direction = Down
	}

}

func (e *Elevator) processDownQueue() {
	e.state.Direction = Down
	e.state.IsMoving = true
	nextRequest := e.requests.downQueue[0]
	if e.state.CurrentFloor == nextRequest.DesiredFloor {
		e.requests.downQueue = e.requests.downQueue[1:]
		e.state.IsMoving = false
		e.state.DoorsOpen = true
		time.Sleep(10 * time.Second)
	} else {
		e.state.CurrentFloor--
		time.Sleep(1 * time.Second)
	}
	if e.state.CurrentFloor == e.minFloor {
		e.state.Direction = Up
	}
}

func (e *Elevator) handleEmergency() {
	e.state.IsMoving = false
	e.state.DoorsOpen = true
	e.state.Direction = Idle

	e.requests.upQueue = make([]Request, 0)
	e.requests.downQueue = make([]Request, 0)
}

func (e *Elevator) ResetEmergency() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.emergency = false
	e.state.DoorsOpen = false
}
