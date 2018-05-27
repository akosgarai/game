package planet

import (
	"github.com/akosgarai/game/pkg/base/speacies"
	"testing"
)

func TestEmpty(t *testing.T) {
	t.Skip("Empty")
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
