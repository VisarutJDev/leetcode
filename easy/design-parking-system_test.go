package easy

import (
	"fmt"
	"testing"
)

func TestAddCar(t *testing.T) {
	parkingSystem := Constructor(1, 1, 0)
	fmt.Println(parkingSystem.AddCar(1))
	fmt.Println(parkingSystem.AddCar(2))
	fmt.Println(parkingSystem.AddCar(3))
	fmt.Println(parkingSystem.AddCar(1))
}
