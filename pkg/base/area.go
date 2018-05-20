package area

import (
    "errors"
    "github.com/akosgarai/game/pkg/base/resource"
    "fmt"
)

var LOG_LEVEL = "ERROR"

type Resource interface {
    Harvest(int) error
    Produce(int)
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
    for index, _ := range a.Resources {
        if a.Resources[index].GetName() == r.GetName() {
            a.Resources[index].Produce(r.GetAmount())
            return nil
        }
    }
    a.Resources = append(a.Resources, r)
    return nil
}

func (a *Area) Build(s Structure) error {
    a.Building = s
    return nil
}

func (a *Area) getResourceByName(name string) (Resource, error) {
    for index, r := range a.Resources {
        if r.GetName() == name {
            return a.Resources[index], nil
        }
    }
    return resource.Empty(), errors.New("Resource missing.")
}

func (a *Area) Harvest () error {

    if a.Building.Name == "" {
        return  nil
    }
    if a.Building.ProducedResource.GetName() == "" {
        return  nil
    }
    if a.Building.NeededResource.GetName() == "" {
        a.AddResource(a.Building.ProducedResource)
        return  nil
    }
    res, err := a.getResourceByName(a.Building.NeededResource.GetName())
    if err != nil {
        return err
    }
    if err = res.Harvest(a.Building.NeededResource.GetAmount()); err != nil {
        return err
    }
    a.AddResource(a.Building.ProducedResource)

    return nil
}
