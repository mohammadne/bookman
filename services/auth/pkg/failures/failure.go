package failures

type Failure interface {
	Message() string
	Status() int
	Causes() []string
	Error() string
}
