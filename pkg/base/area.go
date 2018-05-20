package area

import (
    "errors"
    "github.com/akosgarai/game/pkg/base/resource"
    "fmt"
)

var LOG_LEVEL = "ERROR"

type Resource interface {
    Harvest(int) error
    GetName() string
    GetAmount() int
}

func logger(message string) {
    if LOG_LEVEL == "DEBUG" {
        fmt.Print(message)
    }
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
        return resource.Empty(), nil
    }
    if a.Building.ProducedResource.GetName() == "" {
        return resource.Empty(), nil
    }
    if a.Building.NeededResource.GetName() == "" {
        return a.Building.ProducedResource, nil
    }
    for index, r := range a.Resources {
        if r.GetName() == a.Building.NeededResource.GetName() {
            if err := r.Harvest(a.Building.NeededResource.GetAmount()); err != nil {
                return resource.Empty(), err
            }
            a.Resources[index] = r
            return a.Building.ProducedResource, nil
        }
    }

    return resource.Empty(), errors.New("Resource missing.")
}
