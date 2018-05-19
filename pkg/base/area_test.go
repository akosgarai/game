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
    }
    area.AddResource(r)
    if len(area.Resources) != 1 {
        t.Error("It suppose to be one")
    }
    if area.Resources[0].Name != "goo" {
        t.Error("It suppose to be goo")
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
}
