package miner

import (
	"context"
	"sync"
	"time"
)

type StrongMiner struct {
	MinerInfo
	mtx     sync.RWMutex
	Upgrade int
}

func NewStrongMiner() *StrongMiner {
	return &StrongMiner{
		MinerInfo: MinerInfo{
			Class:       Strong,
			RunsLeft:    60,
			HireCost:    450,
			MaxRuns:     60,
			CoalPerMine: 10,
			Cooldown:    time.Second,
		},
		Upgrade: 3,
	}
}

func (m *StrongMiner) Info() MinerInfo {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.MinerInfo
}

func (m *StrongMiner) Run(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		ticker := time.NewTicker(m.Cooldown)
		defer ticker.Stop()

		for {
			m.mtx.Lock()

			if m.RunsLeft <= 0 {
				m.mtx.Unlock()
				return
			}

			coal := m.CoalPerMine

			m.RunsLeft--
			m.CoalPerMine += m.Upgrade

			m.mtx.Unlock()
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				ch <- coal
			}
		}
	}()

	return ch
}
