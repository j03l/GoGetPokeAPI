package evolution

import (
	"fmt"

	"github.com/j03l/GoGetPokeAPI/pkg/models"
)

// EvolutionChain returns a single evolution chain (by ID).
func EvolutionChain(id string) (result models.EvolutionChain, err error) {
	err = do(fmt.Sprintf("evolution-chain/%s", id), &result)
	return result, err
}

// EvolutionTrigger returns a single evolution trigger (by ID or name).
func EvolutionTrigger(id string) (result models.EvolutionTrigger, err error) {
	err = do(fmt.Sprintf("evolution-trigger/%s", id), &result)
	return result, err
}
