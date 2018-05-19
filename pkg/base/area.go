package area

type Area struct {
    Type string
}

func New () *Area {
    return &Area{
        Type: "grassland",
    }
}
