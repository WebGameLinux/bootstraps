package worker

type Starter interface {
		Init()
		Start()
		AWait() bool
}
