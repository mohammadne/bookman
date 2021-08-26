package network

type Server interface {
	Serve(<-chan struct{})
}
