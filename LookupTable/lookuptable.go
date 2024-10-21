package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Happy Coding @Gophers!")
	lookup := New()
	lookup.Set("Name", "Donna Pysan")
	lookup.Set("12", 24)
	lookup.SetDefaultValue("City", "Macedonia")

	for kv := range lookup.Iterator() {
		fmt.Printf("%s: %v\n", kv.Key, kv.Value)
	}

}

type LookupTable struct {
	data map[any]any
	mu   sync.RWMutex
}

func New() *LookupTable {
	return &LookupTable{
		data: make(map[any]any),
	}
}

func (l *LookupTable) Set(key any, value any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.data[key] = value
}

func (l *LookupTable) Get(key any) (any, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	value, exists := l.data[key]
	return value, exists
}

func (l *LookupTable) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.data)
}

func (l *LookupTable) Delete(key any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.data, key)
}

func (l *LookupTable) Keys() []any {
	l.mu.Lock()
	defer l.mu.Unlock()
	keys := make([]any, 0, len(l.data))
	for k := range l.data {
		keys = append(keys, k)
	}
	return keys
}

func (l *LookupTable) Values() []any {
	l.mu.Lock()
	defer l.mu.Unlock()
	values := make([]any, 0, len(l.data))
	for _, value := range l.data {
		values = append(values, value)
	}
	return values
}

func (l *LookupTable) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.data = make(map[any]any)
}

func (l *LookupTable) Contains(key any) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, exists := l.data[key]
	return exists
}

func (l *LookupTable) Update(other *LookupTable) {
	l.mu.Lock()
	defer l.mu.Unlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for k, v := range other.data {
		l.data[k] = v
	}
}

func (l *LookupTable) Default(key, defaultValue any) any {
	l.mu.Lock()
	defer l.mu.Unlock()
	if value, exists := l.data[key]; exists {
		return value
	}
	return defaultValue
}

func (l *LookupTable) SetDefaultValue(key, defaultValue any) any {
	l.mu.Lock()
	defer l.mu.Unlock()
	if value, exists := l.data[key]; exists {
		return value
	}
	l.data[key] = defaultValue
	return defaultValue
}

func (l *LookupTable) Pop(key any) (any, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if value, exists := l.data[key]; exists {
		delete(l.data, key)
		return value, nil
	}
	return nil, errors.New("key not found")
}

func (l *LookupTable) Iterator() <-chan KeyValue {
	ch := make(chan KeyValue)
	go func() {
		l.mu.RLock()
		defer l.mu.RUnlock()
		for k, v := range l.data {
			ch <- KeyValue{Key: k, Value: v}
		}
		close(ch)
	}()
	return ch
}

type KeyValue struct {
	Key, Value any
}
