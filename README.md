# GoGetPokeAPI

A practice fork of [mtslzr/pokeapi-go](https://github.com/mtslzr/pokeapi-go), modernising the codebase as a learning exercise.

Wrapper for [PokeAPI](https://pokeapi.co), written in Go. *Supports PokeAPI v2.*

## What's changed

- Replacing [patrickmn/go-cache](https://github.com/patrickmn/go-cache) with a custom cache implementation
- Adding structured logging via [charmbracelet/log](https://github.com/charmbracelet/log)
- Modernising project structure and tooling

## Getting Started

```bash
go get github.com/j03l/GoGetPokeAPI
```

```go
import pokeapi "github.com/j03l/GoGetPokeAPI"
```

## Documentation

Full API documentation can be found at [PokeAPI](https://pokeapi.co/docs/v2.html).

## Endpoints

<details>
<summary>Berries</summary>

- Berries, Berry Firmness, Berry Flavors
</details>

<details>
<summary>Contests</summary>

- Contest Types, Contest Effects, Super Contest Effects
</details>

<details>
<summary>Encounters</summary>

- Encounter Methods, Encounter Conditions, Encounter Condition Values
</details>

<details>
<summary>Evolution</summary>

- Evolution Chains, Evolution Triggers
</details>

<details>
<summary>Games</summary>

- Generations, Pokedex, Versions, Version Groups
</details>

<details>
<summary>Items</summary>

- Items, Item Attributes, Item Categories, Item Fling Effects, Item Pockets
</details>

<details>
<summary>Locations</summary>

- Locations, Location Areas, Pal Park Areas, Regions
</details>

<details>
<summary>Machines</summary>

- Machines
</details>

<details>
<summary>Moves</summary>

- Moves, Move Ailments, Move Battle Styles, Move Categories, Move Damage Classes, Move Learn Methods, Move Targets
</details>

<details>
<summary>Pokemon</summary>

- Abilities, Characteristics, Egg Groups, Genders, Growth Rates, Natures, Pokeathlon Stats, Pokemon, Pokemon Colors, Pokemon Forms, Pokemon Habitats, Pokemon Shapes, Pokemon Species, Stats, Types
</details>

<details>
<summary>Utility</summary>

- Languages
</details>

## Acknowledgments

Originally created by [Matthew Salazar](https://github.com/mtslzr/pokeapi-go).
Cache and project architecture inspired by [PokeGo](https://github.com/JoshGuarino/PokeGo) by Josh Guarino.
