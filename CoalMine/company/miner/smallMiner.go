package miner

import (
	"context"
	"sync"
	"time"
)

type SmallMiner struct {
	MinerInfo
	mtx sync.RWMutex
}

func NewSmallMiner() *SmallMiner {
	return &SmallMiner{
		MinerInfo: MinerInfo{
			Class:       Small,
			RunsLeft:    30,
			HireCost:    5,
			MaxRuns:     30,
			CoalPerMine: 1,
			Cooldown:    3 * time.Second,
		},
	}
}

func (m *SmallMiner) Info() MinerInfo {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.MinerInfo
}

func (m *SmallMiner) Run(ctx context.Context) <-chan int {
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
