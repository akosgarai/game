package structure

import (
	"github.com/akosgarai/game/pkg/base/interfaces"
)

type Structure struct {
	name             string
	NeededResource   interfaces.Resource
	ProducedResource interfaces.Resource
}

func Empty() *Structure {
	return &Structure{}
}

func New(name string, needed interfaces.Resource, produced interfaces.Resource) *Structure {
	return &Structure{
		name:             name,
		NeededResource:   needed,
		ProducedResource: produced,
	}
}
func (s *Structure) GetName() string {
	return s.name
}
func (s *Structure) GetNeededResource() interfaces.Resource {
	return s.NeededResource
}
func (s *Structure) GetProducedResource() interfaces.Resource {
	return s.ProducedResource
}
