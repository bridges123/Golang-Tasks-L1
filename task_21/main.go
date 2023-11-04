package main

import "fmt"

// VehicleAdapter определяет методы адаптера
type VehicleAdapter interface {
	Action()
}

// структуры и их методы по умолчанию
type Car struct{}
type Plane struct{}

func (c *Car) Drive(minutes int) {
	fmt.Printf("Car is driving for %d minutes...\n", minutes)
}

func (p *Plane) Fly() {
	fmt.Println("Plane is flying for 60 minutes...")
}

// Адаптеры встраивают реализацию своих структур
type CarAdapter struct {
	*Car
}
type PlaneAdapter struct {
	*Plane
}

// Action реализует метод Drive и заполняет все параметры по умолчанию
func (adapter *CarAdapter) Action() {
	adapter.Drive(60)
}

// Action реализует метод Fly (без параметров)
func (adapter *PlaneAdapter) Action() {
	adapter.Fly()
}

func main() {
	// получаем слайс адаптеров
	vehicles := []VehicleAdapter{&CarAdapter{}, &PlaneAdapter{}}
	// вызываем единый метод для всех
	for _, veh := range vehicles {
		veh.Action()
	}
}
