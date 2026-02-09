package main

import (
    "time"
    "math/rand"
)

type Collector interface {
    Collect() ([]Metric, error)
}

type DummyCollector struct{}

func (c *DummyCollector) Collect() ([]Metric, error) {
    metrics := make([]Metric, 5)
    for i := 0; i < 5; i++ {
        metrics[i] = Metric{
            Name:      "cpu_usage",
            Value:     rand.Float64() * 100,
            Timestamp:   time.Now().Unix(),
            Labels:     map[string]string{"host": "server1", "region": "us-east"},
        }
    }
    return metrics, nil
}
