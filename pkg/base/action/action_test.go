package action

import (
	"github.com/akosgarai/game/pkg/base/resource"
	"testing"
)

func TestAction(t *testing.T) {
	action := Action{
		name:              "testAction",
		isKnown:           false,
		producedResources: resource.New("goo", 10),
		neededResources:   resource.Empty(),
		consumption:       false,
	}

	if action.GetName() != "testAction" {
		t.Error("Invalid action name")
	}
	if action.IsKnown() != false {
		t.Error("Action suppose to be unknown")
	}
	if action.GetConsumption() != false {
		t.Error("Invalid consumption")
	}
	if action.GetNeededResources().GetName() != "" {
		t.Error("Suppose to be empty")
	}
	r := action.GetProducedResources()
	if r.GetName() != "goo" {
		t.Error("Wrong produced resource name")
	}
	if r.GetAmount() != 10 {
		t.Error("Invalid amount")
	}
}
func TestGathering(t *testing.T) {
	action := Gathering()

	if action.GetName() != "Gathering" {
		t.Error("Invalid action name - Gathering")
	}
	if action.IsKnown() != true {
		t.Error("Action suppose to be known")
	}
	if action.GetNeededResources().GetName() != "population" {
		t.Error("Suppose to be population")
	}
	if action.GetNeededResources().GetAmount() != 1 {
		t.Error("Suppose to be 1")
	}
	if action.GetConsumption() != false {
		t.Error("Invalid consumption")
	}
	r := action.GetProducedResources()
	if r.GetName() != "berries" {
		t.Error("Wrong produced resource name")
	}
	if r.GetAmount() != 10 {
		t.Error("Invalid amount")
	}
}
