package event

import (
    "sync"
)
type (
    Storage interface {
        Listeners(string) []Listener
        Add(string, Listener)
        Remove(string, Listener)
    }

    memoryStorage struct {
        rw      sync.RWMutex
        storage map[string][]Listener
    }
)
func NewMemoryStorage() Storage {
    return &memoryStorage{
        storage: make(map[string][]Listener),
    }
}
func (m *memoryStorage) Listeners(s string) []Listener {
    m.rw.RLock()
    defer m.rw.RUnlock()
    ss, _ := m.storage[s]
    return ss
}
func (m *memoryStorage) Add(s string, listener Listener) {
    m.rw.Lock()
    defer m.rw.Unlock()
    listeners, _ := m.storage[s]
    listeners = append(listeners, listener)
    m.storage[s] = listeners
}
func (m *memoryStorage) Remove(s string, listener Listener) {
    m.rw.Lock()
    defer m.rw.Unlock()
    listeners, _ := m.storage[s]

    for i := len(listeners)-1; i >= 0; i-- {
        if listeners[i] == listener {
            listeners = append(listeners[:i], listeners[i+1:]...)
        }
    }
    m.storage[s] = listeners
}
