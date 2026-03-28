package utility

import (
	"fmt"

	"github.com/j03l/GoGetPokeAPI/pkg/models"
)

// Language returns a single language (by name or ID).
func Language(id string) (result models.Language, err error) {
	err = do(fmt.Sprintf("language/%s", id), &result)
	return result, err
}
