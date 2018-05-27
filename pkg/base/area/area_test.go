package area

import (
	"github.com/akosgarai/game/pkg/base/resource"
	"github.com/akosgarai/game/pkg/base/structure"
	"testing"
)

func TestNew(t *testing.T) {
	area := New()
	if area.Type != "grassland" {
		t.Error("It should be grassland!")
	}
	if area.Resources != nil {
		t.Error("It should be nil")
	}
	if area.Building != nil {
		t.Error("It should be nil")
	}
}

func TestAddResource(t *testing.T) {
	area := New()
	if len(area.Resources) != 0 {
		t.Error("It suppose to be empty")
	}

	r := resource.New("goo", 1000)
	area.AddResource(r)
	if len(area.Resources) != 1 {
		t.Error("It suppose to be one")
	}
	if area.Resources[0].GetName() != "goo" {
		t.Error("It suppose to be goo")
	}
	if area.Resources[0].GetAmount() != 1000 {
		t.Error("It suppose to be 1000")
	}
	area.AddResource(r)
	if len(area.Resources) != 1 {
		t.Error("It suppose to be one")
	}
	if area.Resources[0].GetName() != "goo" {
		t.Error("It suppose to be goo")
	}
	if area.Resources[0].GetAmount() != 2000 {
		t.Error("It suppose to be 2000")
	}
}

func TestGetResourceIndexByName(t *testing.T) {
	area := New()
	r := resource.New("goo", 100)
	area.AddResource(r)
	r = resource.New("goo2", 1000)
	area.AddResource(r)
	index, err := area.GetResourceIndexByName("goo")
	if err != nil {
		t.Error("We should find goo.")
	}
	if area.Resources[index].GetName() != "goo" {
		t.Error("We should find goo. We found something else.")
	}
	index, err = area.GetResourceIndexByName("goo2")
	if err != nil {
		t.Error("We should find goo2.")
	}
	if area.Resources[index].GetName() != "goo2" {
		t.Error("We should find goo. We found something else.")
	}
	_, err = area.GetResourceIndexByName("goo3")
	if err == nil {
		t.Error("We shouldn't find goo3.")
	}
	if err.Error() != "Resource missing." {
		t.Error("Different error message for goo3.")
	}
}

func TestBuild(t *testing.T) {
	area := New()
	s := structure.New("Factory", resource.Empty(), resource.Empty())
	area.Build(s)
	if area.Building.GetName() != "Factory" {
		t.Error("It suppose to be Factory")
	}
	s = structure.New("House", resource.Empty(), resource.Empty())
	area.Build(s)
	if area.Building.GetName() != "House" {
		t.Error("It suppose to be House now")
	}
}
func TestAreaHarvest(t *testing.T) {
	area := New()
	err := area.Harvest()
	if err != nil {
		t.Error("Without building we shouldn't have Error")
	}
	building := structure.New("House", resource.Empty(), resource.New("Stuff", 3))
	area.Build(building)
	err = area.Harvest()
	if err != nil {
		t.Error("With this building we shouldn't have Error")
	}
	rsrc_i, rsrcerr := area.GetResourceIndexByName("Stuff")
	rsrc := area.Resources[rsrc_i]
	if rsrcerr != nil {
		t.Error("We should find Stuff without error")
	}
	if rsrc.GetName() != "Stuff" {
		t.Error("With this building we should have Stuff")
	}
	if rsrc.GetAmount() != 3 {
		t.Error("With this building we should have 3 Stuff")
	}
	building = structure.New("Factory", resource.New("Iron", 10), resource.New("Car", 1))
	area.Build(building)
	err = area.Harvest()
	if err == nil {
		t.Error("With this building we should have Error")
	}
	if err.Error() != "Resource missing." {
		t.Error("With this building we should have different Error")
	}
	rsrc_i, rsrcerr = area.GetResourceIndexByName("Car")
	if rsrcerr == nil {
		t.Error("We shouldn't find Car without error")
	}
	new_resource := resource.New("Iron", 1000)
	area.AddResource(new_resource)
	err = area.Harvest()
	if err != nil {
		t.Error("Now, we shouldn't have Error")
	}
	rsrc_i, rsrcerr = area.GetResourceIndexByName("Car")
	rsrc = area.Resources[rsrc_i]
	if rsrcerr != nil {
		t.Error("We should find Car now")
	}
	if rsrc.GetName() != "Car" {
		t.Error("With this building we should have Car")
	}
	if rsrc.GetAmount() != 1 {
		t.Error("With this building we should have 1 Car")
	}
	rsrc_i, rsrcerr = area.GetResourceIndexByName("Iron")
	rsrc = area.Resources[rsrc_i]
	if rsrcerr != nil {
		t.Error("We should find Iron now")
	}
	if rsrc.GetName() != "Iron" {
		t.Error("With this building we should have Iron")
	}
	if rsrc.GetAmount() != 990 {
		t.Errorf("With this building we should have 990 Iron, but we have %d", rsrc.GetAmount())
	}
}
