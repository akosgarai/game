package planet

import(
    "testing"
    "github.com/akosgarai/game/pkg/base/speacies"
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
}
