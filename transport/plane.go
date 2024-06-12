package transport

import "fmt"

type Plane struct {
	Passenger string
}

func (b Plane) ReceivePassengers(passenger string) string {
	fmt.Printf("\nPassanger %v entered Plane", passenger)
	return passenger
}

func (b Plane) DisembarkPassengers(passenger string) string {
	fmt.Printf("\nPassanger %v disembarked from the Plane", passenger)
	return passenger
}
