//Planet is an area container
package planet

import (
    "github.com/akosgarai/game/pkg/base/area"
)

// Planet has Name and size (represented by the number of the Areas)
type Planet struct {
    Name string
    Area []area.Area
}

// New returns *Planet with the given name and size.
func New(name string, size int) *Planet{
    var p = Planet{
        Name: name,
    }
    for i := 0; i < size; i++ {
        p.Area = append(p.Area, *area.New())
    }
    return &p
}
