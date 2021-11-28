package main

import "sync"

type urlQueue struct {
	values []string
	mu     sync.Mutex
	size   int
}

func NewUrlQueue() *urlQueue {
	urlConsumer := urlQueue{}
	urlConsumer.size = 64
	return &urlConsumer
}

func (q *urlQueue) push(url string) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.values) > q.size {
		return false
	}
	q.values = append(q.values, url)
	return true
}

func (q *urlQueue) pop() (bool, string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.values) == 0 {
		return false, ""
	}
	val := q.values[0]
	q.values = q.values[1:]
	return true, val
}

func (q *urlQueue) getSize() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.values)
}
