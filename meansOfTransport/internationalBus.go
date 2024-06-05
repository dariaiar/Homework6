package meansOfTransport

import "fmt"

type Bus struct {
	Passenger string
}

func (i Bus) ReceivePassengers(passenger string) string {
	fmt.Printf("\nPassanger %v entered international Bus", passenger)
	return passenger
}

func (i Bus) DisembarkPassengers(passenger string) string {
	fmt.Printf("\nPassanger %v disembarked from the international Bus", passenger)
	return passenger

}
