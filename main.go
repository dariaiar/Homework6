package main

import (
	"Homework6/meansOfTransport"
	"Homework6/publicTransport"
	"Homework6/routePlans"
	"Homework6/travelRoute"
	"fmt"
)

func main() {
	addTransport, whereToGo := chooseDestination()
	//fmt.Println("\n TEST result addTransport", addTransport)
	fromToBasic := []string{"LocalBus to international terminal"}
	routePlans.MakeRoute(travelRoute.Plan{}, fromToBasic, addTransport)
	passenger := "Petryk Piatochkin"
	WayOfPassenger(whereToGo, passenger)
}

func chooseDestination() ([]string, int) {
	fmt.Printf("\nChoose the route for traveling (print number): 1.Kyiv-Amsterdam; 2.Kyiv-Moon; 3.Kyiv-Athens.\n")
	var whereToGo int
	fmt.Scan(&whereToGo)
	var finish bool
	var addTransport []string
	for !finish {
		switch whereToGo {
		case 1:
			addTransport = []string{"Bus", "Plane", "Train"}
			//fmt.Printf("Following positions were added to the route: %v", addTransport)
			finish = true
		case 2:
			addTransport = []string{"Train", "SpaceShip"}
			//fmt.Printf("Following positions were added to the route: %v", addTransport)
			finish = true
		case 3:
			addTransport = []string{"Bus", "Plane", "Bus"}
			//fmt.Printf("Following positions were added to the route: %v", addTransport)
			finish = true
		default:
			//fmt.Println("Invalid choice. Please choose a valid route number.")
			fmt.Scan(&whereToGo)
		}
	}
	return addTransport, whereToGo
}

func WayOfPassenger(whereToGo int, passenger string) {
	var finish bool
	for !finish {
		switch whereToGo {
		case 1:
			publicTransport.InOutOfTransport(meansOfTransport.LocalBus{}, passenger)
			publicTransport.InOutOfTransport(meansOfTransport.Bus{}, passenger)
			publicTransport.InOutOfTransport(meansOfTransport.Plane{}, passenger)
			publicTransport.InOutOfTransport(meansOfTransport.Train{}, passenger)
			finish = true
		case 2:
			publicTransport.InOutOfTransport(meansOfTransport.LocalBus{}, passenger)
			publicTransport.InOutOfTransport(meansOfTransport.Train{}, passenger)
			publicTransport.InOutOfTransport(meansOfTransport.SpaceShip{}, passenger)
			finish = true
		case 3:

			publicTransport.InOutOfTransport(meansOfTransport.LocalBus{}, passenger)
			publicTransport.InOutOfTransport(meansOfTransport.Bus{}, passenger)
			publicTransport.InOutOfTransport(meansOfTransport.Plane{}, passenger)
			publicTransport.InOutOfTransport(meansOfTransport.Bus{}, passenger)
			finish = true
		default:
			fmt.Println("Something went wrong")

		}
	}

}
