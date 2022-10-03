package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	m sync.Mutex
	v map[string]int
}

func (s *SafeCounter) Incr(key string) {
	s.m.Lock()
	s.v[key]++
	s.m.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.m.Lock()
	defer s.m.Unlock()
	return s.v[key]
}

func main() {
	s := SafeCounter{v: make(map[string]int)}
	s.Incr("mykey")
	s.Incr("mykey")
	s.Incr("mykey")
	fmt.Println(s.v)
	for i:=0;i<10;i++{
		go s.Incr("mykey")
	}
	time.Sleep(time.Second)
	fmt.Println(s.Value("mykey"))
}
