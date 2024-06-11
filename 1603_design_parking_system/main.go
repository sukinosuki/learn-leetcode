package main

type ParkingSystem struct {
	MaxBig    int
	MaxMedium int
	MaxSmall  int

	BigNum    int
	MediumNum int
	SmallNum  int
}

func Constructor(big int, medium int, small int) ParkingSystem {

	return ParkingSystem{
		MaxBig:    big,
		MaxMedium: medium,
		MaxSmall:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {

	if carType == 3 {
		if this.SmallNum == this.MaxSmall {
			return false
		}
		this.SmallNum++
	}

	if carType == 2 {
		if this.MediumNum == this.MaxMedium {
			return false
		}
		this.MediumNum++
	}

	if carType == 1 {
		if this.BigNum == this.MaxBig {
			return false
		}
		this.BigNum++
	}

	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

func main() {

}
