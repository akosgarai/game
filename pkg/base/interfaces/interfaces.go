package interfaces

type Resource interface {
	Harvest(int) error
	Produce(int)
	GetName() string
	GetAmount() int
}

type Building interface {
	GetName() string
	GetNeededResource() Resource
	GetProducedResource() Resource
}

type Action interface {
	IsKnown() bool
	GetName() string
	GetConsumption() bool
	GetNeededResources() Resource
	GetProducedResources() Resource
}
