package pokeapi

type Pokedex struct {
	pokemon map[string]Pokemon
}

type Pokemon struct {
	name string
}

func NewPokedex() Pokedex {
	pokedex := Pokedex{
		pokemon: make(map[string]Pokemon),
	}
	return pokedex
}

func (p *Pokedex) Add(name string) error {
	if _, ok := p.pokemon[name]; !ok {
		p.pokemon[name] = Pokemon{
			name: name,
		}
	}
	return nil
}
