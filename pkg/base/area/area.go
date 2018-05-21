package area

import (
    "errors"
    "github.com/akosgarai/game/pkg/base/interfaces"
    "fmt"
)

var LOG_LEVEL = "ERROR"

func logger(message string) {
    if LOG_LEVEL == "DEBUG" {
        fmt.Print(message)
    }
}

type Building interface {
    GetName() string
    GetNeededResource() interfaces.Resource
    GetProducedResource() interfaces.Resource
}

type Area struct {
    Type string
    Resources []interfaces.Resource
    Building Building
}

func New () *Area {
    return &Area{
        Type: "grassland",
    }
}

func (a *Area) AddResource(r interfaces.Resource) error {
    for index, _ := range a.Resources {
        if a.Resources[index].GetName() == r.GetName() {
            a.Resources[index].Produce(r.GetAmount())
            return nil
        }
    }
    a.Resources = append(a.Resources, r)
    return nil
}

func (a *Area) Build(s Building) error {
    a.Building = s
    return nil
}

func (a *Area) getResourceIndexByName(name string) (int, error) {
    for index, r := range a.Resources {
        if r.GetName() == name {
            return index, nil
        }
    }
    return -1, errors.New("Resource missing.")
}

func (a *Area) Harvest () error {
    if a.Building == nil {
        return nil
    }

    if a.Building.GetName() == "" {
        return  nil
    }
    res := a.Building.GetProducedResource()
    if res.GetName() == "" {
        return  nil
    }
    res_needed := a.Building.GetNeededResource()
    if res_needed.GetName() == "" {
        a.AddResource(a.Building.GetProducedResource())
        return  nil
    }
    res_i, err := a.getResourceIndexByName(res_needed.GetName())
    if err != nil {
        return err
    }
    if err = a.Resources[res_i].Harvest(res_needed.GetAmount()); err != nil {
        return err
    }
    a.AddResource(a.Building.GetProducedResource())

    return nil
}
