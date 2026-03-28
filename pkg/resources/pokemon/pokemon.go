package pokemon

import (
	"fmt"

	"github.com/j03l/GoGetPokeAPI/internal/request"
	"github.com/j03l/GoGetPokeAPI/pkg/models"
)

// Ability returns a single ability (by name or ID).
func Ability(id string) (result models.Ability, err error) {
	err = request.Do(fmt.Sprintf("ability/%s", id), &result)
	return result, err
}

// Characteristic returns a single characteristic (by ID).
func Characteristic(id string) (result models.Characteristic, err error) {
	err = request.Do(fmt.Sprintf("characteristic/%s", id), &result)
	return result, err
}

// EggGroup returns a single egg group (by name or ID).
func EggGroup(id string) (result models.EggGroup, err error) {
	err = request.Do(fmt.Sprintf("egg-group/%s", id), &result)
	return result, err
}

// Gender returns a single gender (by name or ID).
func Gender(id string) (result models.Gender, err error) {
	err = request.Do(fmt.Sprintf("gender/%s", id), &result)
	return result, err
}

// GrowthRate returns a single growth rate (by name or ID).
func GrowthRate(id string) (result models.GrowthRate, err error) {
	err = request.Do(fmt.Sprintf("growth-rate/%s", id), &result)
	return result, err
}

// Nature returns a single nature (by name or ID).
func Nature(id string) (result models.Nature, err error) {
	err = request.Do(fmt.Sprintf("nature/%s", id), &result)
	return result, err
}

// PokeathlonStat returns a single Pokeathlon state (by name or ID).
func PokeathlonStat(id string) (result models.PokeathlonStat, err error) {
	err = request.Do(fmt.Sprintf("pokeathlon-stat/%s", id), &result)
	return result, err
}

// Pokemon returns a single Pokemon (by name or ID).
func Pokemon(id string) (result models.Pokemon, err error) {
	err = request.Do(fmt.Sprintf("pokemon/%s", id), &result)
	return result, err
}

// PokemonColor returns a single Pokemon color (by name or ID).
func PokemonColor(id string) (result models.PokemonColor, err error) {
	err = request.Do(fmt.Sprintf("pokemon-color/%s", id), &result)
	return result, err
}

// PokemonForm returns a single Pokemon form (by name or ID).
func PokemonForm(id string) (result models.PokemonForm, err error) {
	err = request.Do(fmt.Sprintf("pokemon-form/%s", id), &result)
	return result, err
}

// PokemonHabitat returns a single Pokemon habitat (by name or ID).
func PokemonHabitat(id string) (result models.PokemonHabitat, err error) {
	err = request.Do(fmt.Sprintf("pokemon-habitat/%s", id), &result)
	return result, err
}

// PokemonShape returns a single Pokemon shape (by name or ID).
func PokemonShape(id string) (result models.PokemonShape, err error) {
	err = request.Do(fmt.Sprintf("pokemon-shape/%s", id), &result)
	return result, err
}

// PokemonSpecies returns a single Pokemon species (by name or ID).
func PokemonSpecies(id string) (result models.PokemonSpecies, err error) {
	err = request.Do(fmt.Sprintf("pokemon-species/%s", id), &result)
	return result, err
}

// Stat returns a single stat (by name or ID).
func Stat(id string) (result models.Stat, err error) {
	err = request.Do(fmt.Sprintf("stat/%s", id), &result)
	return result, err
}

// Type returns a single type (by name or ID).
func Type(id string) (result models.Type, err error) {
	err = request.Do(fmt.Sprintf("type/%s", id), &result)
	return result, err
}
