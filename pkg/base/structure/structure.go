package structure

type Resource interface {
    Harvest(int) error
    Produce(int)
    GetName() string
    GetAmount() int
}

type Structure struct {
    name string
    NeededResource Resource
    ProducedResource Resource
}

func Empty() *Structure {
    return &Structure{}
}

func New (name string, needed Resource, produced Resource) *Structure {
    return &Structure{
        name: name,
        NeededResource: needed,
        ProducedResource: produced,
    }
}
func (s *Structure) GetName() string {
    return s.name
}
func (s *Structure) GetNeededResource() Resource {
    return s.NeededResource
}
func (s *Structure) GetProducedResource() Resource {
    return s.ProducedResource
}
