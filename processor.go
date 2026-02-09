package main

import (
    "sync"
    "time"
)

type Processor struct {
    queue   *Queue
    storage map[string][]AggregatedMetric
    mu      sync.Mutex
}

func NewProcessor(queue *Queue) *Processor {
    return &Processor{
        queue:   queue,
        storage: make(map[string][]AggregatedMetric),
    }
}

func (p *Processor) Process() {
    metrics := p.queue.PopAll()
    if len(metrics) == 0 {
        return
    }

    // Группируем по имени метрики
    grouped := make(map[string][]Metric)
    for _, m := range metrics {
        grouped[m.Name] = append(group packed[m.Name], m)
    }

    for name, metrics := range grouped {
        agg := AggregatedMetric{
            Name:      name,
            AvgValue:  0,
            MinValue:  metrics[0].Value,
            MaxValue:  metrics[0].Value,
            Count:     len(metrics),
            PeriodStart: metrics[0].Timestamp,
            PeriodEnd:   metrics[len(metrics)-1].Timestamp,
        }

        var sum float64
        for _, m := range metrics {
            sum += m.Value
            if m.Value < agg.MinValue {
                agg.MinValue = m.Value
            }
            if m.Value > agg.MaxValue {
                agg.MaxValue = m.Value
            }
        }
        agg.AvgValue = sum / float64(len(metrics))

        p.mu.Lock()
        p.storage[name] = append(p.storage[name], agg)
        p.mu.Unlock()
    }
}
