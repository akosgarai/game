package interfaces

type Resource interface {
	Harvest(int) error
	Produce(int)
	GetName() string
	GetAmount() int
}
