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
- Type - grassland
- Resources - list of resources, that are availeable on the area
- Building - it's a kind of structure. It will be used to transform basic resources to better stuff.

### planet

it contains X amount of areas depending on the planet size

### population

TBD

### technology

if invented, it can give some improvements, or something cool stuff, it can open other technologies
