package storage

import "time"

type Storage struct{}

func NewStorage(historySize time.Duration) *Storage {
	return &Storage{}
}
