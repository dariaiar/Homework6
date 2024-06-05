package meansOfTransport

import "fmt"

type SpaceShip struct {
	Passenger string
}

func (b SpaceShip) ReceivePassengers(passenger string) string {
	fmt.Printf("\nPassanger %v entered Space Ship", passenger)
	return passenger
}

func (b SpaceShip) DisembarkPassengers(passenger string) string {
	fmt.Printf("\nPassanger %v disembarked from the Space Ship", passenger)
	return passenger

}
