package web

type Server interface {
	Serve(<-chan struct{}) error
}

type Client interface{}
