package multicastChannel

import "sync"

type TypedMulticast[T any] struct {
	mu         sync.RWMutex
	channelCap int
	listeners  []chan T
	closed     bool
}

func NewTypedMulticast[T any](channelCap int) *TypedMulticast[T] {
	return &TypedMulticast[T]{
		channelCap: channelCap,
		listeners:  make([]chan T, 0),
	}
}

func (m *TypedMulticast[T]) Subscribe() <-chan T {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed {
		return nil
	}

	ch := make(chan T, m.channelCap)
	m.listeners = append(m.listeners, ch)
	return ch
}

func (m *TypedMulticast[T]) Publish(msg T) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.closed {
		return
	}

	for _, listener := range m.listeners {
		select {
		case listener <- msg:
		default:
			// 处理满的 channel
		}
	}
}

func (m *TypedMulticast[T]) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.closed = true
	for _, listener := range m.listeners {
		close(listener)
	}
	m.listeners = nil
}
