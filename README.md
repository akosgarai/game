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

### Structure

Constructed stuff.

- Name

### area

It reperesents the single unit of area. Each area has type.

```golang
type Area struct {
    Type string
    Resources []Resource
    Building Structure
}
```

- Type - grassland
- Resources - list of resources, that are availeable on the area
- Building - it's a kind of structure. It will be used to transform basic resources to better stuff.
- AddResource(Resource) error - It increases the number of our Resource.Name amount with Resource.Amount or inserts the new Resource to the Resources.
- Build(Structure) error - It replaces the current area.Building with the given Structure.
- getResourceByName(string) (Resource, error) - Returns the Resource object if we have Resource.Name resource, or returns error if we haven't.
- Harvest() error - Modifies the amount of resources, according to the Building.

### planet

it contains X amount of areas depending on the planet size

### population

TBD

### technology

if invented, it can give some improvements, or something cool stuff, it can open other technologies
