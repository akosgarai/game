// Action is a kind of transformer
package action

import (
	"github.com/akosgarai/game/pkg/base/interfaces"
	"github.com/akosgarai/game/pkg/base/resource"
)

// Action is something that needs some resources, and produces other ones.
// name - Name of the action.
// isknown - is true, if the action is default or invented.
// neededResources - The name and the amount of the resource, that is needed for the action.
//The actions are activietes, by default, so the needed resources are not used.
// producedResources - The resource that is constructed during the action.
type Action struct {
	name              string
	neededResources   interfaces.Resource
	producedResources interfaces.Resource
	isKnown           bool
}

// IsKnown - returns true, if the action is known.
func (a *Action) IsNkown() bool {
	return a.isKnown
}

// GetName returns the action's name
func (a *Action) GetName() string {
	return a.name
}

// GetNeededResources returns the needes resources
func (a *Action) GetNeededResources() interfaces.Resource {
	return a.neededResources
}

// GetProducedResources returns the constructed resource
func (a *Action) GetProducedResources() interfaces.Resource {
	return a.producedResources
}

// Gathering returns the gathering action
func Gathering() *Action {
	return &Action{
		name:              "Gathering",
		isKnown:           true,
		producedResources: resource.New("berries", 10),
		neededResources:   resource.New("population", 1),
	}
}
