package transport

import "fmt"

type LocalBus struct {
	Passenger string
}

func (b LocalBus) ReceivePassengers(passenger string) string {
	fmt.Printf("\nPassanger %v entered Local Bus to reach international terminal", passenger)
	return passenger
}

func (b LocalBus) DisembarkPassengers(passenger string) string {
	fmt.Printf("\nPassanger %v disembarked from Local Bus", passenger)
	return passenger
}
