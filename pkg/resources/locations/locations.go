package locations

import (
	"fmt"

	"github.com/j03l/GoGetPokeAPI/internal/request"
	"github.com/j03l/GoGetPokeAPI/pkg/models"
)

// Location returns a single location (by name or ID).
func Location(id string) (result models.Location, err error) {
	err = request.Do(fmt.Sprintf("location/%s", id), &result)
	return result, err
}

// LocationArea returns a single location area (by name or ID).
func LocationArea(id string) (result models.LocationArea, err error) {
	err = request.Do(fmt.Sprintf("location-area/%s", id), &result)
	return result, err
}

// PalParkArea returns a single Pal Park area (by name or ID).
func PalParkArea(id string) (result models.PalParkArea, err error) {
	err = request.Do(fmt.Sprintf("pal-park-area/%s", id), &result)
	return result, err
}

// Region returns a single region (by name or ID).
func Region(id string) (result models.Region, err error) {
	err = request.Do(fmt.Sprintf("region/%s", id), &result)
	return result, err
}
