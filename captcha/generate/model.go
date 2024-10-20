package generate

import (
	"crypto/rand"
	"fmt"
	"sync"
	"time"
)

type CaptchaStore struct {
	store map[string]string
	mu    sync.Mutex
}

func NewCaptchaStore() *CaptchaStore {
	return &CaptchaStore{
		store: make(map[string]string),
	}
}

func (cs *CaptchaStore) Set(id, value string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.store[id] = value
	go func() {
		time.Sleep(2 * time.Minute)
		cs.Delete(id)
	}()
}

func (cs *CaptchaStore) Get(id string) (string, bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	value, exists := cs.store[id]
	return value, exists
}

func (cs *CaptchaStore) Delete(id string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	delete(cs.store, id)
}

func generateCaptcha(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
