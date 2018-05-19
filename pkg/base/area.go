package area

type Resource struct {
    Name string
}

type Structure struct {
    Name string
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
