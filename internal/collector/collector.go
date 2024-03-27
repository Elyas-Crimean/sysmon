package collector

import "github.com/Elyas-Crimean/sysmon/internal/storage"

type Collector struct {
	storage *storage.Storage
}

func NewCollector(storage *storage.Storage) *Collector {
	return &Collector{storage: storage}
}

func (c *Collector) Run() {}
