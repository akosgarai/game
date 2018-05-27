# The game project

I want to implement a cool planet/solar system/galaxy management application.

## Structures

### pkg/base/resource

It represents something useful stuff that the population should use.

```golang
type Resource struct {
    Name string
    Amount int
}
```

- Name - name of the resource.
- Amount - the amount of the resource
- Harvest(int) error - decrease the amount of the resource with the given input.
- Produce(int) - increase the amount of the resource with the given input.
- GetName() string - return the name of the resource
- GetAmount() int - returns the amount of the resource

### pkg/base/speacies

Resource implementation for representing the population.

```golang
type Speacies struct {
    Name string
    population int
}
```

- Name - name of the speacies
- population - the number of the population
- Harvest(int) error - decrease the number of the population with the given input.
- Produce(int) - increase the number of the population with the given input.
- GetName() string - return the name of the speacies
- GetAmount() int - returns the number of the population

### pkg/base/structure

Constructed stuff.

```golang
type Structure struct {
    name string
    NeededResource *resource.Resource
    ProducedResource *resource.Resource
}
```

- name - name of the structure
- NeededResource - stuff that needs to be presented to be able to build new stuff.
- ProducedResource - stuffs that will be produced, if the NeededResources are given.
- GetName() string - return the name of the structure
- GetNeededResource() interfaces.Resource - return the Needed resource
- GetProducedResource() interfaces.Resource - return the produced resource

### pkg/base/area

It reperesents the single unit of area. Each area has type.

```golang
type Area struct {
    Type string
    Resources []resource.Resource
    Building Building
}
```

- Type - grassland
- Resources - list of resources, that are availeable on the area
- Building - it's a kind of structure. It will be used to transform basic resources to better stuff.
- AddResource(interfaces.Resource) error - It increases the number of our Resource.Name amount with Resource.Amount or inserts the new Resource to the Resources.
- Build(Building) error - It replaces the current area.Building with the given Structure.
- getResourceIndexByName(string) (int, error) - Returns the index of Resource object (in resources) if we have Resource.Name resource, or returns error if we haven't.
- Harvest() error - Modifies the amount of resources, according to the Building.

### planet

it contains X amount of areas depending on the planet size

### population

TBD

### technology

if invented, it can give some improvements, or something cool stuff, it can open other technologies
