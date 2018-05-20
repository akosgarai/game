package area

import (
    "errors"
    "strconv"
    "fmt"
)

var LOG_LEVEL = "ERROR"

type Resource struct {
    Name string
    Amount int
}

func logger(message string) {
    if LOG_LEVEL == "DEBUG" {
        fmt.Print(message)
    }
}

func (r *Resource) Harvest(amount int) error {
    message := "Harvesting from resource '" + r.Name + "' " + strconv.Itoa(amount) + " / " + strconv.Itoa(r.Amount) + "\n"
    logger(message)
    if r.Amount < amount {
        return errors.New("Insufficient number of resources.")
    }
    r.Amount = r.Amount - amount
    message = "Harvesting from resource '" + r.Name + "' " + strconv.Itoa(amount) + " / " + strconv.Itoa(r.Amount) + "\n"
    logger(message)
    return nil
}

type Structure struct {
    Name string
    NeededResource Resource
    ProducedResource Resource
}

type Area struct {
    Type string
    Resources []Resource
    Building Structure
}

func New () *Area {
    return &Area{
        Type: "grassland",
    }
}

func (a *Area) AddResource(r Resource) error {
    a.Resources = append(a.Resources, r)
    return nil
}

func (a *Area) Build(s Structure) error {
    a.Building = s
    return nil
}

func (a *Area) Harvest () (Resource, error) {

    if a.Building.Name == "" {
        return Resource{}, nil
    }
    if a.Building.ProducedResource.Name == "" {
        return Resource{}, nil
    }
    if a.Building.NeededResource.Name == "" {
        return a.Building.ProducedResource, nil
    }
    for index, r := range a.Resources {
        if r.Name == a.Building.NeededResource.Name {
            if err := r.Harvest(a.Building.NeededResource.Amount); err != nil {
                return Resource{}, err
            }
            a.Resources[index] = r
            return a.Building.ProducedResource, nil
        }
    }

    return Resource{}, errors.New("Resource missing.")
}
