//Resource implementation for representing the population.
package speacies

//Speacies is a kind of resource.
//It means, it has name (Name), and amount (population).
type Speacies struct {
    Name string
    population int
}

// New Returns *Speacies with the given inputs (Name and initial population)
func New(name string, initialPopulation int) *Speacies {
    return &Speacies{
        Name: name,
        population: initialPopulation,
    }
}

// Produce increase the amount of the resource with the given input.
func (s *Speacies) Produce(base int) {
    s.population = s.population + base
}

// Harvest decrease the amount of the resource with the given input.
// Returns error - mostly when the population goes below zero. It means extinction.
func (s *Speacies) Harvest(base int) error {
    s.population = s.population - base
    return nil
}

// GetName returns the Name of the speacies.
func (s *Speacies) GetName() string {
    return s.Name
}

// GetAmount returns the number of the population.
func (s *Speacies) GetAmount() int {
    return s.population
}
