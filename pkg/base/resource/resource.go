package resource

import(
    "errors"
    "strconv"
)

type Resource struct {
    Name string
    Amount int
}

func New(name string, amount int) *Resource {
    return &Resource{
        Name: name,
        Amount: amount,
    }
}

func (r *Resource) Harvest(amount int) error {
    if r.Amount < amount {
        message := "Insufficient. Need: " + strconv.Itoa(amount) + ", Amount: " + strconv.Itoa(r.Amount)
        return errors.New(message)
    }
    r.Amount = r.Amount - amount
    return nil
}
