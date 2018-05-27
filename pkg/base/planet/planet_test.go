package planet

import (
	"github.com/akosgarai/game/pkg/base/action"
	"github.com/akosgarai/game/pkg/base/interfaces"
	"github.com/akosgarai/game/pkg/base/resource"
	"github.com/akosgarai/game/pkg/base/speacies"
	"testing"
)

type TestAction struct {
	known       bool
	consumption bool
}

func (t *TestAction) IsKnown() bool {
	return t.known
}
func (t *TestAction) GetConsumption() bool {
	return false
}
func (t *TestAction) GetName() string {
	return "TestAction"
}
func (t *TestAction) GetNeededResources() interfaces.Resource {
	return resource.New("population", 100)
}
func (t *TestAction) GetProducedResources() interfaces.Resource {
	return resource.Empty()
}

func TestNew(t *testing.T) {
	p := New("TestName", 3)
	if p.Name != "TestName" {
		t.Error("Invalid name for plaet")
	}
	if len(p.Area) != 3 {
		t.Error("Area length invalid")
	}
}

func TestPopulate(t *testing.T) {
	p := New("TestName", 3)
	guys := speacies.New("test race", 1)
	err := p.Populate(guys)
	if err != nil {
		t.Error("Population suppose to be successful first")
	}
	err = p.Populate(guys)
	if err == nil {
		t.Error("Population suppose to be errored next time")
	}
	if err.Error() != "Planet populated." {
		t.Error("Different error message")
	}
	if len(p.Actions) != 1 {
		t.Error("Should have action after population")
	}
	if p.Actions[0].GetName() != "Gathering" {
		t.Error("Should have gathering action after population")
	}
}

func TestAddAction(t *testing.T) {
	p := New("TestName", 3)
	guys := speacies.New("test race", 1)
	err := p.Populate(guys)
	if err != nil {
		t.Error("Population suppose to be successful first")
	}
	if len(p.Actions) != 1 {
		t.Error("Should have action after population")
	}
	if p.Actions[0].GetName() != "Gathering" {
		t.Error("Should have gathering action after population")
	}
	p.addAction(action.Gathering())
	if len(p.Actions) != 1 {
		t.Error("Acton's length suppose to be the same")
	}
}
func TestGetAction(t *testing.T) {
	p := New("TestName", 3)
	guys := speacies.New("test race", 1)
	err := p.Populate(guys)
	if err != nil {
		t.Error("Population suppose to be successful first")
	}
	actioni, err := p.getActionIndex("Gathering")
	if err != nil {
		t.Error("Gathering suppose to be known")
	}
	if p.Actions[actioni].GetName() != "Gathering" {
		t.Error("Gathering suppose to be named as Gathering")
	}
	_, err = p.getActionIndex("FutureTech")
	if err == nil {
		t.Error("FutureTech suppose to be unknown")
	}
	otherAction := &TestAction{
		known: false,
	}
	p.addAction(otherAction)
	_, err = p.getActionIndex("TestAction")
	if err == nil {
		t.Error("TestAction suppose to be unknown")
	}
}
func TestDoActionGathering(t *testing.T) {
	p := New("TestName", 3)
	guys := speacies.New("test race", 1)
	err := p.Populate(guys)
	if err != nil {
		t.Error("Population suppose to be successful first")
	}
	err = p.DoAction("Riot")
	if err == nil {
		t.Error("Riot suppose to be unsuccessful")
	}
	err = p.DoAction("Gathering")
	if err != nil {
		t.Error("Gathering suppose to be successful")
	}
	index, err := p.Area[0].GetResourceIndexByName("berries")
	if err != nil {
		t.Error(err.Error())
	}
	if index == -1 {
		t.Error("Index out of range")
	}
	if p.Area[0].Resources[index].GetAmount() != 10 {
		t.Error("Unexpected amount")
	}
	otherAction := &TestAction{
		known: true,
	}
	p.addAction(otherAction)
	err = p.DoAction("TestAction")
	if err == nil {
		t.Error("Suppose to get error")
	}
	if err.Error() != "Not enough population for doing that action." {
		t.Error("Suppose to get different error message")
	}

}
