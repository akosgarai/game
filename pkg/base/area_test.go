package area

import(
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
    if area.Building.Name != "" {
        t.Error("It should be empty")
    }
}

func TestAddResource(t *testing.T) {
    area := New()
    if len(area.Resources) != 0 {
        t.Error("It suppose to be empty")
    }

    r := Resource{
        Name : "goo",
        Amount: 1000,
    }
    area.AddResource(r)
    if len(area.Resources) != 1 {
        t.Error("It suppose to be one")
    }
    if area.Resources[0].Name != "goo" {
        t.Error("It suppose to be goo")
    }
    if area.Resources[0].Amount != 1000 {
        t.Error("It suppose to be 1000")
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
func TestResourceHarvest(t *testing.T) {
    r := Resource{
        Name : "goo",
        Amount: 1000,
    }
    r.Harvest(10)
    if r.Amount != 990 {
        t.Error("It suppose to be 990")
    }
    r.Harvest(10)
    if r.Amount != 980 {
        t.Error("It suppose to be 980")
    }
    r.Harvest(10)
    if r.Amount != 970 {
        t.Error("It suppose to be 970")
    }
    err := r.Harvest(1000)
    if err == nil {
        t.Error("It suppose to be error")
    }
}
func TestAreaHarvest(t *testing.T) {
    area := New()
    resource, err := area.Harvest()
    if err != nil {
        t.Error("Without building we shouldn't have Error")
    }
    if resource.Name != "" {
        t.Error("Without building we shouldn't have Resource")
    }
    building := Structure{
        Name: "House",
        NeededResource: Resource{},
        ProducedResource: Resource{
            Name: "Stuff",
            Amount: 3,
        },
    }
    area.Build(building)
    resource, err = area.Harvest()
    if err != nil {
        t.Error("With this building we shouldn't have Error")
    }
    if resource.Name != "Stuff" {
        t.Error("With this building we should have Stuff")
    }
    if resource.Amount != 3 {
        t.Error("With this building we should have 3 Stuff")
    }
    building = Structure{
        Name: "Factory",
        NeededResource: Resource{
            Name: "Iron",
            Amount: 10,
        },
        ProducedResource: Resource{
            Name: "Car",
            Amount: 1,
        },
    }
    area.Build(building)
    resource, err = area.Harvest()
    if err == nil {
        t.Error("With this building we should have Error")
    }
    if err.Error() != "Resource missing." {
        t.Error("With this building we should have different Error")
    }
    if resource.Name != "" {
        t.Error("With this building we shouldn't have name")
    }
    new_resource := Resource{
        Name: "Iron",
        Amount: 1000,
    }
    area.AddResource(new_resource)
    resource, err = area.Harvest()
    if err != nil {
        t.Error("Now, we shouldn't have Error")
    }
    if resource.Name != "Car" {
        t.Error("With this building we should have Car")
    }
    if resource.Amount != 1 {
        t.Error("With this building we should have 1 Car")
    }
    for _, r := range area.Resources {
        if r.Name == "Iron" {
            if r.Amount != 990 {
                t.Errorf("This building should use 10 resource for it's output. Amount: %d", r.Amount)
            }
        }
    }
}
