package bitcoinvipsvr

import "sync"

type syncQueue struct {
	lock    *sync.RWMutex
	element []interface{}
}

func newQueue() *syncQueue {
	return &syncQueue{
		lock: new(sync.RWMutex),
	}
}

func (q *syncQueue) Push(e interface{}) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	q.element = append(q.element, e)
}

func (q *syncQueue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	firstElement := q.element[0]
	q.element = q.element[1:]
	return firstElement
}

func (q *syncQueue) Clear() bool {
	if q.IsEmpty() {
		return false
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	len := len(q.element)
	for i := 0; i < len; i++ {
		q.element[i] = nil
	}
	q.element = nil
	return true
}

func (q *syncQueue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if len(q.element) == 0 {
		return true
	}
	return false
}

func (q *syncQueue) size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.element)
}
