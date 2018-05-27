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

// addAction insert new action to the list
// it has some basic validation.
func (p *Planet) addAction(a interfaces.Action) {
	name := a.GetName()
	if name == "" || !p.isActionInvented(name) {
		p.Actions = append(p.Actions, a)
	}
}

// isActionInvented is a helper function to prevent multiple action insertion
func (p *Planet) isActionInvented(name string) bool {
	for _, v := range p.Actions {
		if v.GetName() == name {
			return true
		}
	}
	return false
}

// getActionIndex returns the index of the action and nil.
// On error, returns -1, error
func (p *Planet) getActionIndex(name string) (int, error) {
	if !p.isActionInvented(name) {
		return -1, errors.New("Action isn't invented")
	}
	for i, v := range p.Actions {
		if v.GetName() == name && v.IsKnown() {
			return i, nil
		}
	}
	return -1, errors.New("Action isn't invented")
}

// DoAction - should do an action given by it's name.
// Returns nil or error
func (p *Planet) DoAction(name string) error {
	index, err := p.getActionIndex(name)
	if err != nil {
		return err
	}
	action := p.Actions[index]
	switch action.GetNeededResources().GetName() {
	case "population":
		if action.GetNeededResources().GetAmount() > p.Population.GetAmount() {
			return errors.New("Not enough population for doing that action.")
		}
		if action.GetConsumption() {
			err = p.Population.Harvest(action.GetNeededResources().GetAmount())
			if err != nil {
				return err
			}
		}
		err = p.Area[0].AddResource(action.GetProducedResources())
		if err != nil {
			return err
		}
		break
	default:
		break
	}
	return nil
}
