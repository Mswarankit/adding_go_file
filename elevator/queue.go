package elevator

func (r *RequestQueue) addUpRequest(req Request) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	inserted := false

	for i, rq := range r.upQueue {
		if req.DesiredFloor < rq.DesiredFloor {
			r.upQueue = append(r.upQueue[:i], append([]Request{req}, r.upQueue[i:]...)...)
			inserted = true
			break
		}
		if !inserted {
			r.upQueue = append(r.upQueue, req)
		}
	}
}

func (r *RequestQueue) addDownRequest(req Request) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	inserted := false

	for i, rq := range r.downQueue {
		if req.DesiredFloor > rq.DesiredFloor {
			r.downQueue = append(r.downQueue[:i], append([]Request{req}, r.downQueue[i:]...)...)
			inserted = true
			break
		}
		if !inserted {
			r.downQueue = append(r.downQueue, req)
		}
	}
}
