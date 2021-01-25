package easy

// Design a parking system for a parking lot. The parking lot has three kinds of parking spaces: big, medium, and small, with a fixed number of slots for each size.
// Implement the ParkingSystem class:
// 		ParkingSystem(int big, int medium, int small) Initializes object of the ParkingSystem class.
//		The number of slots for each parking space are given as part of the constructor.
//		bool addCar(int carType) Checks whether there is a parking space of carType for the car that wants to get into the parking lot.
//		carType can be of three kinds: big, medium, or small, which are represented by 1, 2, and 3 respectively.
//		A car can only park in a parking space of its carType. If there is no space available, return false, else park the car in that size space and return true.

// ParkingSystem struct
type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

// Constructor fundtion
func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

// AddCar a function to add car
func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.Big == 0 {
			return false
		}
		this.Big--
		return true
	case 2:
		if this.Medium == 0 {
			return false
		}
		this.Medium--
		return true
	case 3:
		if this.Small == 0 {
			return false
		}
		this.Small--
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
