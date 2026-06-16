package company

var EquipmentCatalog = map[string]int{
	"Pickaxe":     3000,
	"Ventilation": 15000,
	"Carts":       50000,
}

type Equipment struct {
	Name   string
	Cost   int
	Bought bool
}
