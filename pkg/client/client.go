package messagebroker

// Client contains necessary values for interacting with the message broker.
type client struct {
	uri      string
	username string
	password string
}

type Raz struct {
	name string
}

// Client creates and returns a client for interacting with the message broker.
func Connect(name string) string {
	return "razee"
}
