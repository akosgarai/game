package area

type Resource struct {
    Name string
}

type Area struct {
    Type string
    Resources []Resource
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
