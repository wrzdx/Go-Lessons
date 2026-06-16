package server

import (
	"CoalMine/company"
	"encoding/json"
	"time"
)

type StartDTO struct {
	Message string
	Started time.Time
}

type FinishDTO struct {
	Message string
	company.CompanyStats
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
