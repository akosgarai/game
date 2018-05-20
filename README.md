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
- GetName() string - return the name of the resource
- GetAmount() int - returns the amount of the resource

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
- GetNeededResource() string - return the name of the Needed resource
- GetProducedResource() string - return the name of the produced resource

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
- AddResource(\*resource.Resource) error - It increases the number of our Resource.Name amount with Resource.Amount or inserts the new Resource to the Resources.
- Build(Building) error - It replaces the current area.Building with the given Structure.
- getResourceByName(string) (resource.Resource, error) - Returns the Resource object if we have Resource.Name resource, or returns error if we haven't.
- Harvest() error - Modifies the amount of resources, according to the Building.

### pkg/base/planet

it contains X amount of areas depending on the planet size

```golang
type Planet struct {
    Name string
    Area []area.Area
}
```

- Name - Name of the planet
- Area - Size of the planet - 1 area / size.

### population

TBD

### technology

if invented, it can give some improvements, or something cool stuff, it can open other technologies
