package company

import (
	"CoalMine/company/miner"
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

type Company struct {
	Started   *time.Time
	Finished  *time.Time
	Balance   atomic.Int64
	ctx       context.Context
	cancel    context.CancelFunc
	Staff     []miner.Miner
	Equipment map[string]bool
}

type CompanyStats struct {
	Elapsed   string
	Staff     []miner.MinerInfo
	Balance   int
	Equipment []Equipment
}

func NewCompany() *Company {
	ctx, cancel := context.WithCancel(context.Background())
	equipment := make(map[string]bool, len(EquipmentCatalog))
	for name := range EquipmentCatalog {
		equipment[name] = false
	}
	return &Company{
		ctx:       ctx,
		cancel:    cancel,
		Equipment: equipment,
	}
}

func (c *Company) Start() error {
	if c.Started != nil {
		return ErrAlreadyStarted
	}
	now := time.Now()
	c.Started = &now
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-c.ctx.Done():
				return

			case <-ticker.C:
				c.Balance.Add(1)
			}
		}
	}()

	return nil
}

func (c *Company) BuyEquipment(name string) error {
	if c.Started == nil {
		return ErrNotStarted
	}
	if c.Finished != nil {
		return ErrAlreadyFinished
	}
	cost, ok := EquipmentCatalog[name]
	if !ok {
		return ErrEquipmentNotFound
	}
	if c.Equipment[name] {
		return ErrAlreadyBought
	}

	if cost > int(c.Balance.Load()) {
		return ErrNotEnoughMoney
	}

	c.Balance.Add(-int64(cost))
	c.Equipment[name] = true

	return nil
}

func (c *Company) GotAllEquipment() bool {
	for _, isBought := range c.Equipment {
		if !isBought {
			return false
		}
	}
	return true
}

func (c *Company) Finish() error {
	if c.Started == nil {
		return ErrNotStarted
	}
	if c.Finished != nil {
		return ErrAlreadyFinished
	}
	if !c.GotAllEquipment() {
		return ErrEquipmentNoBought
	}
	c.cancel()
	now := time.Now()
	c.Finished = &now

	return nil
}

func (c *Company) HireMiner(class string) (miner.MinerInfo, error) {
	if c.Started == nil {
		return miner.MinerInfo{}, ErrNotStarted
	}
	if c.Finished != nil {
		return miner.MinerInfo{}, ErrAlreadyFinished
	}
	factory, ok := miner.MinerFactories[class]
	if !ok {
		return miner.MinerInfo{}, ErrMinerNotFound
	}

	newMiner := factory()
	cost := int64(newMiner.Info().HireCost)

	if c.Balance.Load() < cost {
		return miner.MinerInfo{}, ErrNotEnoughMoney
	}
	c.Balance.Add(-cost)
	c.Staff = append(c.Staff, newMiner)
	go func() {
		ch := newMiner.Run(c.ctx)
		for coal := range ch {
			c.Balance.Add(int64(coal))
		}
	}()

	return newMiner.Info(), nil
}

func (c *Company) StaffInfo() []miner.MinerInfo {
	staff := make([]miner.MinerInfo, 0, len(c.Staff))

	for _, m := range c.Staff {
		staff = append(staff, m.Info())
	}

	return staff
}

func (c *Company) ActiveStaffInfo() []miner.MinerInfo {
	allStaff := c.StaffInfo()

	active := make([]miner.MinerInfo, 0, len(allStaff))

	for _, m := range allStaff {
		if m.RunsLeft > 0 {
			active = append(active, m)
		}
	}

	return active
}

func (c *Company) EquipmentInfo() []Equipment {
	equipment := make([]Equipment, 0, len(c.Equipment))

	for name, bought := range c.Equipment {
		equipment = append(equipment, Equipment{Name: name, Cost: EquipmentCatalog[name], Bought: bought})
	}

	return equipment
}

func (c *Company) Stats() CompanyStats {
	stats := CompanyStats{
		Balance:   int(c.Balance.Load()),
		Staff:     c.StaffInfo(),
		Equipment: c.EquipmentInfo(),
	}

	if c.Started != nil && c.Finished != nil {
		elapsed := c.Finished.Sub(*c.Started)
		stats.Elapsed = fmt.Sprintf("%02d:%02d:%02d", int(elapsed.Hours()), int(elapsed.Minutes())%60, int(elapsed.Seconds())%60)
	}

	return stats
}
