package main

func (p *Processor) GetAggregatedMetrics(name string) []AggregatedMetric {
    p.mu.Lock()
    metrics := p.storage[name]
    p.mu.Unlock()
    return metrics
}
