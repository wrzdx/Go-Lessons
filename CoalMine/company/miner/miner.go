package miner

import (
	"context"
	"time"
)

const (
	Small  = "Small"
	Normal = "Normal"
	Strong = "Strong"
)

type MinerInfo struct {
	Class       string
	RunsLeft    int
	HireCost    int
	MaxRuns     int
	CoalPerMine int
	Cooldown    time.Duration
}

type Miner interface {
	Run(ctx context.Context) <-chan int
	Info() MinerInfo
}

var MinerFactories = map[string]func() Miner{
	Small: func() Miner {
		return NewSmallMiner()
	},
	Normal: func() Miner {
		return NewNormalMiner()
	},
	Strong: func() Miner {
		return NewStrongMiner()
	},
}

func GetMinersInfo() []MinerInfo {
	minersInfo := make([]MinerInfo, 0, len(MinerFactories))
	for _, factory := range MinerFactories {
		minersInfo = append(minersInfo, factory().Info())
	}

	return minersInfo
}
