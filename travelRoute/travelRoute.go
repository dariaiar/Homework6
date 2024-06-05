package travelRoute

import "fmt"

type Plan struct {
	RouteUnif []string
}

func (b Plan) AddTransportToRoutes(fromToBasic []string, addTransport []string) []string {
	fromToBasic = append(fromToBasic, b.RouteUnif...)
	fromToFinal := append(fromToBasic, addTransport...)
	fmt.Println("\nAdditional transport was added to the route:", addTransport)
	return fromToFinal
}

func (b Plan) ShowTransportOnTheRout(fromToFinal []string) []string {
	fmt.Println("\nFinal list of transport on the route is", fromToFinal)
	return fromToFinal
}
