package resource

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	r := Empty()
	if r.GetName() != "" {
		t.Errorf("Name suppose to be empty. We have: name: %s, amount: %d", r.GetName(), r.GetAmount())
	}
}

func TestNew(t *testing.T) {

	var testData = []struct {
		Name   string
		Amount int
	}{
		{"test1", 10},
		{"test2", 20},
	}

	for _, tt := range testData {
		resource := New(tt.Name, tt.Amount)
		if resource.GetName() != tt.Name {
			t.Errorf("Name suppose to be same. %s instead of %s", resource.Name, tt.Name)
		}
		if resource.GetAmount() != tt.Amount {
			t.Errorf("Amount suppose to be same. %d instead of %d", resource.Amount, tt.Amount)
		}
	}
}

func TestHarvest(t *testing.T) {
	r := Resource{
		Name:   "goo",
		Amount: 1000,
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
	err := r.Harvest(1000)
	if err == nil {
		t.Error("It suppose to be error")
	}
	expected_error := "Insufficient. Need: 1000, Amount: 970"
	if err.Error() != expected_error {
		t.Error("Unexpected error message")
	}
}
func TestProduce(t *testing.T) {
	r := Resource{
		Name:   "goo",
		Amount: 10,
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
