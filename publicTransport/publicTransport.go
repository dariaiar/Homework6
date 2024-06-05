package publicTransport

type Transport interface {
	ReceivePassengers(string) string
	DisembarkPassengers(string) string
}

func InOutOfTransport(t Transport, passenger string) {
	t.ReceivePassengers(passenger)
	t.DisembarkPassengers(passenger)
}
