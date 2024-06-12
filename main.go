package main

import (
	"Homework6/transport"
	"Homework6/travel"
	"fmt"
)

func main() {
	addTransport, whereToGo := chooseDestination()
	//fmt.Println("\n TEST result addTransport", addTransport)
	fromToBasic := []string{"LocalBus to international terminal"}
	travel.MakeRoute(travel.Plan{}, fromToBasic, addTransport)
	passenger := "Petryk Piatochkin"
	WayOfPassenger(whereToGo, passenger)
}

func chooseDestination() ([]string, int) {
	fmt.Printf("\nChoose the route for traveling (print number): 1.Kyiv-Amsterdam; 2.Kyiv-Moon; 3.Kyiv-Athens.\n")
	var whereToGo int
	for {
		fmt.Scan(&whereToGo)
		switch whereToGo {
		case 1:
			return []string{"Bus", "Plane", "Train"}, 1
		case 2:
			return []string{"Train", "SpaceShip"}, 2
		case 3:
			return []string{"Bus", "Plane", "Bus"}, 3
		default:
			fmt.Println("Unknown route, try again!")
		}
	}
}

func WayOfPassenger(whereToGo int, passenger string) {
	var finish bool
	for !finish {
		switch whereToGo {
		case 1:
			travel.InOutOfTransport(transport.LocalBus{}, passenger)
			travel.InOutOfTransport(transport.Bus{}, passenger)
			travel.InOutOfTransport(transport.Plane{}, passenger)
			travel.InOutOfTransport(transport.Train{}, passenger)
			finish = true
		case 2:
			travel.InOutOfTransport(transport.LocalBus{}, passenger)
			travel.InOutOfTransport(transport.Train{}, passenger)
			travel.InOutOfTransport(transport.SpaceShip{}, passenger)
			finish = true
		case 3:
			travel.InOutOfTransport(transport.LocalBus{}, passenger)
			travel.InOutOfTransport(transport.Bus{}, passenger)
			travel.InOutOfTransport(transport.Plane{}, passenger)
			travel.InOutOfTransport(transport.Bus{}, passenger)
			finish = true
		default:
			fmt.Println("Something went wrong")
		}
	}
}
