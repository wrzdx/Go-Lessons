package miner

import (
	"context"
	"sync"
	"time"
)

type NormalMiner struct {
	MinerInfo
	mtx sync.RWMutex
}

func NewNormalMiner() *NormalMiner {
	return &NormalMiner{
		MinerInfo: MinerInfo{
			Class:       Normal,
			RunsLeft:    45,
			HireCost:    50,
			MaxRuns:     45,
			CoalPerMine: 3,
			Cooldown:    2 * time.Second,
		},
	}
}

func (m *NormalMiner) Info() MinerInfo {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.MinerInfo
}

func (m *NormalMiner) Run(ctx context.Context) <-chan int {
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
