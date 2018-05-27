// Simulator package. Only for trying out long term solutions
package main

import (
	"fmt"
	"github.com/akosgarai/game/pkg/base/planet"
	"github.com/akosgarai/game/pkg/base/speacies"
)

func main() {
	Planet := planet.New("Simulus Romus", 5)
	Habitants := speacies.New("Rabbit-sapiens", 1)
	Planet.Populate(Habitants)
	fmt.Print("Planet ", Planet.Name, " has been populated with ", Habitants.GetName())
}
