package games

import (
	"fmt"

	"github.com/j03l/GoGetPokeAPI/internal/request"
	"github.com/j03l/GoGetPokeAPI/pkg/models"
)

// Generation returns a single generation (by name or ID).
func Generation(id string) (result models.Generation, err error) {
	err = request.Do(fmt.Sprintf("generation/%s", id), &result)
	return result, err
}

// Pokedex returns a single Pokedex (by name or ID).
func Pokedex(id string) (result models.Pokedex, err error) {
	err = request.Do(fmt.Sprintf("pokedex/%s", id), &result)
	return result, err
}

// Version returns a single version (by name or ID).
func Version(id string) (result models.Version, err error) {
	err = request.Do(fmt.Sprintf("version/%s", id), &result)
	return result, err
}

// VersionGroup returns a single version group (by name or ID).
func VersionGroup(id string) (result models.VersionGroup, err error) {
	err = request.Do(fmt.Sprintf("version-group/%s", id), &result)
	return result, err
}
