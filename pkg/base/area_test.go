package area

import(
    "testing"
)

func TestNew(t *testing.T) {
    area := New()
    if area.Type != "grassland" {
        t.Error("It should be grassland!")
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
