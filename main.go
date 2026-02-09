package main

import (
    "log"
    "time"
)

func main() {
    queue := NewQueue()
    collector := &DummyCollector{}
    processor := NewProcessor(queue)

    processor.RegisterHandlers()

    // Запуск сборщика каждые 10 секунд
    go func() {
        for {
            metrics, err := collector.Collect()
            if err != nil {
                log.Printf("Error collecting metrics: %v", err)
                continue
            }
            for _, m := range metrics {
                queue.Push(m)
            }
            time.Sleep(10 * time.Second)
        }
    }()

    // Обработка каждые 15 секунд
    go func() {
        for {
            processor.Process()
            time.Sleep(15 * time.Second)
        }
    }()

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
