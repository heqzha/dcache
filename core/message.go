package core

import (
	"sync"

	"github.com/heqzha/goutils/container"
)

type MessageQueue struct {
	q     map[string]*container.Queue
	mutex *sync.RWMutex
}

func (m *MessageQueue) Init() {
	m.q = make(map[string]*container.Queue)
	m.mutex = &sync.RWMutex{}
}

func (m *MessageQueue) Push(chn string, msg interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.q[chn]
	if !ok {
		m.q[chn] = new(container.Queue)
	}
	m.q[chn].Push(msg)
}

func (m *MessageQueue) Pop(chn string) interface{} {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	c, ok := m.q[chn]
	if ok && c != nil {
		return c.Pop()
	}
	return nil
}

func (m *MessageQueue) Len(chn string) int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.q[chn].Len()
}
