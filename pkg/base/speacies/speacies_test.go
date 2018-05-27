package speacies

import (
	"testing"
)

func TestNew(t *testing.T) {

	var testData = []struct {
		Name       string
		population int
	}{
		{"test1", 10},
		{"test2", 20},
	}

	for _, tt := range testData {
		resource := New(tt.Name, tt.population)
		if resource.GetName() != tt.Name {
			t.Errorf("Name suppose to be same. %s instead of %s", resource.Name, tt.Name)
		}
		if resource.GetAmount() != tt.population {
			t.Errorf("population suppose to be same. %d instead of %d", resource.population, tt.population)
		}
	}
}

func TestHarvest(t *testing.T) {
	r := Speacies{
		Name:       "goo",
		population: 1000,
	}
	r.Harvest(10)
	if r.GetAmount() != 990 {
		t.Error("It suppose to be 990")
	}
	r.Harvest(10)
	if r.GetAmount() != 980 {
		t.Error("It suppose to be 980")
	}
	r.Harvest(10)
	if r.GetAmount() != 970 {
		t.Error("It suppose to be 970")
	}
}
func TestProduce(t *testing.T) {
	r := Speacies{
		Name:       "goo",
		population: 10,
	}
	r.Produce(10)
	if r.GetAmount() != 20 {
		t.Error("It suppose to be 20")
	}
	r.Produce(10)
	if r.GetAmount() != 30 {
		t.Error("It suppose to be 30")
	}
	r.Produce(10)
	if r.GetAmount() != 40 {
		t.Error("It suppose to be 40")
	}
}
