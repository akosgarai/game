//Planet is an area container
package planet

import (
	"errors"
	"github.com/akosgarai/game/pkg/base/action"
	"github.com/akosgarai/game/pkg/base/area"
	"github.com/akosgarai/game/pkg/base/interfaces"
)

// Planet has Name and size (represented by the number of the Areas)
// Name - Name of the planet.
// Area - Size of the planet - 1 area / size.
// Population - the speacies that lives in the planet.
// isPopulated - true if the planet is populated, Population is set and not extincted.
type Planet struct {
	Name        string
	Area        []area.Area
	Population  interfaces.Resource
	isPopulated bool
	Actions     []interfaces.Action
}

// New returns *Planet with the given name and size.
func New(name string, size int) *Planet {
	var p = Planet{
		Name:        name,
		isPopulated: false,
	}
	for i := 0; i < size; i++ {
		p.Area = append(p.Area, *area.New())
	}
	return &p
}

// Populate add Population resource to the planet.
func (p *Planet) Populate(r interfaces.Resource) error {
	if p.isPopulated {
		return errors.New("Planet populated.")
	}
	p.Population = r
	p.addAction(action.Gathering())
	p.isPopulated = true
	return nil
}

func (p *Planet) addAction(a interfaces.Action) {
	p.Actions = append(p.Actions, a)
}
