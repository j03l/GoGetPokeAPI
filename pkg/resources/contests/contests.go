package contests

import (
	"fmt"

	"github.com/j03l/GoGetPokeAPI/pkg/models"
)

// ContestType returns a single contest type (by name or ID).
func ContestType(id string) (result models.ContestType, err error) {
	err = do(fmt.Sprintf("contest-type/%s", id), &result)
	return result, err
}

// ContestEffect returns a single contest effect (by ID).
func ContestEffect(id string) (result models.ContestEffect, err error) {
	err = do(fmt.Sprintf("contest-effect/%s", id), &result)
	return result, err
}

// SuperContestEffect returns a single super contest effect (by ID).
func SuperContestEffect(id string) (result models.SuperContestEffect, err error) {
	err = do(fmt.Sprintf("super-contest-effect/%s", id), &result)
	return result, err
}
