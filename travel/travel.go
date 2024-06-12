package travel

import "fmt"

type Transport interface {
	ReceivePassengers(string) string
	DisembarkPassengers(string) string
}

func InOutOfTransport(t Transport, passenger string) {
	t.ReceivePassengers(passenger)
	t.DisembarkPassengers(passenger)
}

type Route interface {
	AddTransportToRoutes([]string, []string) []string
	ShowTransportOnTheRout([]string) []string
}

func MakeRoute(r Route, fromToBasic []string, addTransport []string) []string {
	fmt.Println("\nBasic list of transport for all routs is:", fromToBasic)
	fromToFinal := r.AddTransportToRoutes(fromToBasic, addTransport)
	//fmt.Println("\nFinal list of transport on the route:", fromToFinal)
	r.ShowTransportOnTheRout(fromToFinal)
	return fromToFinal
}

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
