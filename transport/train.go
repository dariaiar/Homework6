package transport

import "fmt"

type Train struct {
	Passenger string
}

func (b Train) ReceivePassengers(passenger string) string {
	fmt.Printf("\nPassanger %v entered Train", passenger)
	return passenger
}

func (b Train) DisembarkPassengers(passenger string) string {
	fmt.Printf("\nPassanger %v disembarked from the Train", passenger)
	return passenger

}
