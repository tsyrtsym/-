package main

import "sync"

type Queue struct {
    items []Metric
    mu    sync.Mutex
}

func NewQueue() *Queue {
    return &Queue{
        items: make([]Metric, 0),
    }
}

func (q *Queue) Push(metric Metric) {
    q.mu.Lock()
    q.items = append(q.items, metric)
    q.mu.Unlock()
}

func (q *Queue) PopAll() []Metric {
    q.mu.Lock()
    items := q.items
    q.items = make([]Metric, 0)
    q.mu.Unlock()
    return items
}
