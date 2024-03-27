package storage

import "time"

type Storage struct {
	limit time.Duration
}

func NewStorage(historySize time.Duration) *Storage {
	return &Storage{limit: historySize}
}

func (s Storage) GetAvg(window time.Duration) (result map[string]float64) {
	_ = window
	return
}
