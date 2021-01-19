package models

var (
	CoffeeMachineList map[string]*CoffeeMachine
)

type CoffeeMachine struct {
	ProductType     ProductType
	WaterLine		bool
}

type ProductType int

const (
	COFFEE_MACHINE_LARGE ProductType = iota
	COFFEE_MACHINE_SMALL
	ESPRESSO_MACHINE
)

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func init() {
	CoffeeMachineList = make(map[string]*CoffeeMachine)
}

func GetAllCoffeeMachines() map[string]*CoffeeMachine {
	return CoffeeMachineList
}
