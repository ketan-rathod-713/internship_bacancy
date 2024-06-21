package stringer

import "fmt"

type DistanceUnit int

const (
	Kilometer DistanceUnit = 0
	Mile      DistanceUnit = 1
)

type Distance struct {
	number float64
	unit   DistanceUnit
}

func (u DistanceUnit) String() string {
	if u == Kilometer {
		return "km"
	}

	if u == Mile {
		return "mile"
	}
	return "unit"
}

func (d Distance) String() string { // It must be pass by value for stringer ha ha
	return fmt.Sprintf("%v %v", d.number, d.unit)
}

func InterfaceStringerExample3() {
	d := Distance{number: 30, unit: Kilometer} // before String it will print {30 0}
	fmt.Println(d)

	e := Distance{number: 50, unit: 20}
	fmt.Println(e)
}

// Now lets add string method for Distance unit and also for distance
