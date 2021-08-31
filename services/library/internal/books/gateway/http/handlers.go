package http

type Handlers interface{}

type handlersImpl struct{}

func New() Handlers {
	return &handlersImpl{}
}
