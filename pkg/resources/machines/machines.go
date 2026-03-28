package machines

import (
	"fmt"

	"github.com/j03l/GoGetPokeAPI/pkg/models"
)

// Machine returns a single machine (by ID).
func Machine(id string) (result models.Machine, err error) {
	err = do(fmt.Sprintf("machine/%s", id), &result)
	return result, err
}
