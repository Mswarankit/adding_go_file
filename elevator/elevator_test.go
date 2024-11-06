package elevator

import (
	"sync"
	"testing"
	"time"
)

func TestNewElevator(t *testing.T) {
	elevator := NewElevator(1, 45)
	if elevator.state.CurrentFloor != 1 {
		t.Errorf("Expected initial floor to be 1, got %d", elevator.state.CurrentFloor)
	}
}

func TestElevatorQueueProcessing(t *testing.T) {
	elevator := NewElevator(1, 10)

	// Add multiple requests
	requests := []Request{
		{CurrentFloor: 1, DesiredFloor: 5, IsInternal: true},
		{CurrentFloor: 1, DesiredFloor: 3, IsInternal: true},
		{CurrentFloor: 5, DesiredFloor: 2, IsInternal: true},
	}

	for _, req := range requests {
		err := elevator.AddRequest(req)
		if err != nil {
			t.Fatalf("Failed to add request: %v", err)
		}
	}

	elevator.Start()
	time.Sleep(1 * time.Second)

	status := elevator.GetStatus()
	if status.Direction == Up {
		t.Error("Elevator should be moving")
	}
}

// elevator/elevator_test.go

func TestProcessDownQueue(t *testing.T) {
	tests := []struct {
		name          string
		initialFloor  int
		requests      []Request
		expectedFloor int
		expectedDir   Direction
		waitTime      time.Duration
	}{
		{
			name:         "Single floor down movement",
			initialFloor: 5,
			requests: []Request{
				{CurrentFloor: 5, DesiredFloor: 3, IsInternal: true},
			},
			expectedFloor: 4,
			expectedDir:   Down,
			waitTime:      2 * time.Second,
		},
		{
			name:         "Multiple floor down requests",
			initialFloor: 10,
			requests: []Request{
				{CurrentFloor: 10, DesiredFloor: 8, IsInternal: true},
				{CurrentFloor: 8, DesiredFloor: 6, IsInternal: true},
				{CurrentFloor: 6, DesiredFloor: 4, IsInternal: true},
			},
			expectedFloor: 8,
			expectedDir:   Down,
			waitTime:      3 * time.Second,
		},
		{
			name:         "Reach bottom floor",
			initialFloor: 3,
			requests: []Request{
				{CurrentFloor: 3, DesiredFloor: 1, IsInternal: true},
			},
			expectedFloor: 1,
			expectedDir:   Up, // Should change direction at bottom
			waitTime:      4 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elevator := NewElevator(1, 10)
			elevator.state.CurrentFloor = tt.initialFloor

			for _, req := range tt.requests {
				err := elevator.AddRequest(req)
				if err != nil {
					t.Fatalf("Failed to add request: %v", err)
				}
			}

			elevator.Start()
			time.Sleep(tt.waitTime)

			status := elevator.GetStatus()
			if status.CurrentFloor != tt.expectedFloor {
				t.Errorf("Expected floor %d, got %d", tt.expectedFloor, status.CurrentFloor)
			}
			if status.Direction != tt.expectedDir {
				t.Errorf("Expected direction %v, got %v", tt.expectedDir, status.Direction)
			}
		})
	}
}

func TestEmergencyHandling(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(*Elevator)
		checkFunc func(*testing.T, *Elevator)
		waitTime  time.Duration
	}{
		{
			name: "Emergency while moving",
			setupFunc: func(e *Elevator) {
				e.AddRequest(Request{CurrentFloor: 1, DesiredFloor: 5, IsInternal: true})
				e.Start()
				time.Sleep(500 * time.Millisecond)
				e.TriggerEmergency()
			},
			checkFunc: func(t *testing.T, e *Elevator) {
				status := e.GetStatus()
				if !status.DoorsOpen {
					t.Error("Doors should be open during emergency")
				}
				if status.IsMoving {
					t.Error("Elevator should not be moving during emergency")
				}
				if status.Direction != Idle {
					t.Error("Elevator should be in Idle state during emergency")
				}
				if len(e.requests.upQueue) > 0 || len(e.requests.downQueue) > 0 {
					t.Error("Request queues should be cleared during emergency")
				}
			},
			waitTime: 1 * time.Second,
		},
		{
			name: "Emergency reset",
			setupFunc: func(e *Elevator) {
				e.TriggerEmergency()
				time.Sleep(500 * time.Millisecond)
				e.ResetEmergency()
			},
			checkFunc: func(t *testing.T, e *Elevator) {
				status := e.GetStatus()
				if status.DoorsOpen {
					t.Error("Doors should be closed after emergency reset")
				}
				if e.emergency {
					t.Error("Emergency flag should be false after reset")
				}
			},
			waitTime: 1 * time.Second,
		},
		{
			name: "Request handling during emergency",
			setupFunc: func(e *Elevator) {
				e.TriggerEmergency()
				e.AddRequest(Request{CurrentFloor: 1, DesiredFloor: 5, IsInternal: true})
			},
			checkFunc: func(t *testing.T, e *Elevator) {
				if len(e.requests.upQueue) > 0 || len(e.requests.downQueue) > 0 {
					t.Error("Requests should not be processed during emergency")
				}
			},
			waitTime: 500 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elevator := NewElevator(1, 10)

			if tt.setupFunc != nil {
				tt.setupFunc(elevator)
			}

			time.Sleep(tt.waitTime)

			if tt.checkFunc != nil {
				tt.checkFunc(t, elevator)
			}
		})
	}
}

func TestEmergencyEdgeCases(t *testing.T) {
	t.Run("Multiple emergency triggers", func(t *testing.T) {
		elevator := NewElevator(1, 10)

		// Trigger emergency multiple times
		for i := 0; i < 3; i++ {
			elevator.TriggerEmergency()
			status := elevator.GetStatus()
			if !status.DoorsOpen {
				t.Error("Doors should remain open with multiple emergency triggers")
			}
		}
	})

	t.Run("Emergency during door operation", func(t *testing.T) {
		elevator := NewElevator(1, 10)

		// Add request and start elevator
		elevator.AddRequest(Request{CurrentFloor: 1, DesiredFloor: 2, IsInternal: true})
		elevator.Start()

		// Wait for elevator to start moving
		time.Sleep(500 * time.Millisecond)

		// Trigger emergency
		elevator.TriggerEmergency()

		status := elevator.GetStatus()
		if !status.DoorsOpen {
			t.Error("Doors should open immediately when emergency is triggered")
		}
		if status.IsMoving {
			t.Error("Elevator should stop immediately when emergency is triggered")
		}
	})
}

// Helper function to test concurrent emergency handling
func TestConcurrentEmergencyHandling(t *testing.T) {
	elevator := NewElevator(1, 10)
	elevator.Start()

	// Create multiple goroutines to simulate concurrent emergency situations
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			elevator.TriggerEmergency()
			time.Sleep(100 * time.Millisecond)
			elevator.ResetEmergency()
		}()
	}

	wg.Wait()

	// Verify final state
	status := elevator.GetStatus()
	if status.IsMoving {
		t.Error("Elevator should not be moving after concurrent emergency handling")
	}
}
