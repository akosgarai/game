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

type Building interface {
    GetName() string
    GetNeededResource() *resource.Resource
    GetProducedResource() *resource.Resource
}

type Area struct {
    Type string
    Resources []resource.Resource
    Building Building
}

func New () *Area {
    return &Area{
        Type: "grassland",
    }
}

func (a *Area) AddResource(r *resource.Resource) error {
    for index, _ := range a.Resources {
        if a.Resources[index].GetName() == r.GetName() {
            a.Resources[index].Produce(r.GetAmount())
            return nil
        }
    }
    a.Resources = append(a.Resources, *r)
    return nil
}

func (a *Area) Build(s Building) error {
    a.Building = s
    return nil
}

func (a *Area) getResourceByName(name string) (*resource.Resource, error) {
    for index, r := range a.Resources {
        if r.GetName() == name {
            return &a.Resources[index], nil
        }
    }
    return resource.Empty(), errors.New("Resource missing.")
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
    res, err := a.getResourceByName(res_needed.GetName())
    if err != nil {
        return err
    }
    if err = res.Harvest(res_needed.GetAmount()); err != nil {
        return err
    }
    a.AddResource(a.Building.GetProducedResource())

    return nil
}
