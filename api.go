package main

import (
    "encoding/json"
    "net/http"
)

func (p *Processor) RegisterHandlers() {
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        if name == "" {
            http.Error(w, "name parameter required", http.StatusBadRequest)
            return
        }

        metrics := p.GetAggregatedMetrics(name)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(metrics)
    })
}
