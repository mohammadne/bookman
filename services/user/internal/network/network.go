package network

type Server interface {
	Serve(<-chan struct{})
}

type Client interface {
	Setup()
}
