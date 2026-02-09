package main

type Metric struct {
    Name      string            `json:"name"`
    Value     float64           `json:"value"`
    Timestamp int64             `json:"timestamp"`
    Labels    map[string]string `json:"labels,omitempty"`
}

type AggregatedMetric struct {
    Name       string  `json:"name"`
    AvgValue   float64 `json:"avg_value"`
    MinValue   float64 `json:"min_value"`
    MaxValue   float64 `json:"max_value"`
    Count      int     `json:"count"`
    PeriodStart int64  `json:"period_start"`
    PeriodEnd   int64  `json:"period_end"`
}
