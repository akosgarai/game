package structure

import(
    "github.com/akosgarai/game/pkg/base/resource"
)

type Structure struct {
    name string
    NeededResource resource.Resource
    ProducedResource resource.Resource
}

func Empty() *Structure {
    return &Structure{}
}

func New (name string, needed resource.Resource, produced resource.Resource) *Structure {
    return &Structure{
        name: name,
        NeededResource: needed,
        ProducedResource: produced,
    }
}
func (s *Structure) GetName() string {
    return s.name
}
func (s *Structure) GetNeededResource() resource.Resource {
    return s.NeededResource
}
func (s *Structure) GetProducedResource() resource.Resource {
    return s.ProducedResource
}
