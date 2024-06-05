package routePlans

import "fmt"

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
