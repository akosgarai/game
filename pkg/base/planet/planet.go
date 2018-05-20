package planet

import (
    "github.com/akosgarai/game/pkg/base/area"
)

type Planet struct {
    Name string
    Area []area.Area
}

func New(name string, size int) *Planet{
    var p = Planet{
        Name: name,
    }
    for i := 0; i < size; i++ {
        p.Area = append(p.Area, *area.New())
    }
    return &p
}
