package company

import "errors"

var ErrAlreadyStarted = errors.New("Company already started")
var ErrAlreadyFinished = errors.New("Company already finished")
var ErrNotStarted = errors.New("Company has not started")
var ErrEquipmentNoBought = errors.New("Company has not bought all equipment")
var ErrEquipmentNotFound = errors.New("Equipment is not found")
var ErrNotEnoughMoney = errors.New("Company has not enough money to buy this equipment/miner")
var ErrMinerNotFound = errors.New("Miner of this class not found")
var ErrAlreadyBought = errors.New("This equipment already bought")
