package main

import "sync"

var exists = struct{}{}

type webIndexSet struct {
	m           map[string]struct{}
	readCounter int
	mu          sync.Mutex
}

func NewWebIndexSet() *webIndexSet {
	s := &webIndexSet{}
	s.m = make(map[string]struct{})
	return s
}

func (s *webIndexSet) Add(value string) {
	s.mu.Lock()
	s.m[value] = exists
	s.mu.Unlock()
}

func (s *webIndexSet) Remove(value string) {
	s.mu.Lock()
	delete(s.m, value)
	s.mu.Unlock()
}

func (s *webIndexSet) Contains(value string) bool {
	s.mu.Lock()
	_, c := s.m[value]
	defer s.mu.Unlock()
	return c
}
