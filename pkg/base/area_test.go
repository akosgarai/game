package area

import(
    "testing"
    "github.com/akosgarai/game/pkg/base/resource"
)

func TestNew(t *testing.T) {
    area := New()
    if area.Type != "grassland" {
        t.Error("It should be grassland!")
    }
    if area.Resources != nil {
        t.Error("It should be nil")
    }
    if area.Building.Name != "" {
        t.Error("It should be empty")
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

func TestGetResourceByName(t *testing.T) {
    area := New()
    r := resource.New("goo", 100)
    area.AddResource(r)
    r = resource.New("goo2", 1000)
    area.AddResource(r)
    _, err := area.getResourceByName("goo")
    if err != nil {
        t.Error("We should find goo.")
    }
    _, err = area.getResourceByName("goo2")
    if err != nil {
        t.Error("We should find goo2.")
    }
    _, err = area.getResourceByName("goo3")
    if err == nil {
        t.Error("We shouldn't find goo3.")
    }
    if err.Error() != "Resource missing." {
        t.Error("Different error message for goo3.")
    }
}

func TestBuild(t *testing.T) {
    area := New()
    s := Structure{
        Name : "Factory",
    }
    area.Build(s)
    if area.Building.Name != "Factory" {
        t.Error("It suppose to be Factory")
    }
    s = Structure{
        Name: "House",
    }
    area.Build(s)
    if area.Building.Name != "House" {
        t.Error("It suppose to be House now")
    }
}
func TestAreaHarvest(t *testing.T) {
    area := New()
    err := area.Harvest()
    if err != nil {
        t.Error("Without building we shouldn't have Error")
    }
    building := Structure{
        Name: "House",
        NeededResource: resource.Empty(),
        ProducedResource: resource.New("Stuff", 3),
    }
    area.Build(building)
    err = area.Harvest()
    if err != nil {
        t.Error("With this building we shouldn't have Error")
    }
    rsrc, rsrcerr := area.getResourceByName("Stuff")
    if rsrcerr != nil {
        t.Error("We should find Stuff without error")
    }
    if rsrc.GetName() != "Stuff" {
        t.Error("With this building we should have Stuff")
    }
    if rsrc.GetAmount() != 3 {
        t.Error("With this building we should have 3 Stuff")
    }
    building = Structure{
        Name: "Factory",
        NeededResource: resource.New("Iron", 10),
        ProducedResource: resource.New("Car", 1),
    }
    area.Build(building)
    err = area.Harvest()
    if err == nil {
        t.Error("With this building we should have Error")
    }
    if err.Error() != "Resource missing." {
        t.Error("With this building we should have different Error")
    }
    rsrc, rsrcerr = area.getResourceByName("Car")
    if rsrcerr == nil {
        t.Error("We shouldn't find Car without error")
    }
    new_resource := resource.New("Iron", 1000)
    area.AddResource(new_resource)
    err = area.Harvest()
    if err != nil {
        t.Error("Now, we shouldn't have Error")
    }
    rsrc, rsrcerr = area.getResourceByName("Car")
    if rsrcerr != nil {
        t.Error("We should find Car now")
    }
    if rsrc.GetName() != "Car" {
        t.Error("With this building we should have Car")
    }
    if rsrc.GetAmount() != 1 {
        t.Error("With this building we should have 1 Car")
    }
    rsrc, rsrcerr = area.getResourceByName("Iron")
    if rsrcerr != nil {
        t.Error("We should find Iron now")
    }
    if rsrc.GetName() != "Iron" {
        t.Error("With this building we should have Iron")
    }
    if rsrc.GetAmount() != 990 {
        t.Error("With this building we should have 990 Iron")
    }
}
