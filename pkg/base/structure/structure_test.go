package structure

import(
    "testing"
    "github.com/akosgarai/game/pkg/base/resource"
)

func TestEmpty(t *testing.T) {
    s := Empty()
    if s.GetName() != "" {
        t.Errorf("Name suppose to be empty. We have: name: %s", s.name)
    }
}

func TestNew(t *testing.T) {
    t.Skip("New")
}

func TestGetName(t *testing.T) {
    s := Empty()
    if s.GetName() != "" {
        t.Errorf("Name suppose to be empty. We have: name: %s", s.name)
    }
    s = New("TestName", resource.Empty(), resource.Empty())
    if s.GetName() != "TestName" {
        t.Errorf("Name suppose to be TestName. We have: name: %s", s.name)
    }
}

func TestGetNeededResource(t *testing.T) {
    s := New("TestName", resource.Empty(), resource.Empty())
    res := s.GetNeededResource()
    if res.GetName() != "" {
        t.Errorf("We shouldn't get name. We have %s", res.GetName())
    }
    s = New("TestName2", resource.New("Iron", 10), resource.New("Car", 1))
    res = s.GetNeededResource()
    if res.GetName() != "Iron" {
        t.Errorf("It should be Iron. We have %s", res.GetName())
    }
}

func TestGetProducedResource(t *testing.T) {
    s := New("TestName", resource.Empty(), resource.Empty())
    res := s.GetProducedResource()
    if res.GetName() != "" {
        t.Errorf("We shouldn't get name. We have %s", res.GetName())
    }
    s = New("TestName2", resource.New("Iron", 10), resource.New("Car", 1))
    res = s.GetProducedResource()
    if res.GetName() != "Car" {
        t.Errorf("It should be Car. We have %s", res.GetName())
    }
}

