package encounters

import (
	"fmt"

	"github.com/j03l/GoGetPokeAPI/internal/request"
	"github.com/j03l/GoGetPokeAPI/pkg/models"
)

// EncounterMethod returns a single encounter method (by name or ID).
func EncounterMethod(id string) (result models.EncounterMethod, err error) {
	err = request.Do(fmt.Sprintf("encounter-method/%s", id), &result)
	return result, err
}

// EncounterCondition returns a single encounter condition (by name or ID).
func EncounterCondition(id string) (result models.EncounterCondition,
	err error) {
	err = request.Do(fmt.Sprintf("encounter-condition/%s", id), &result)
	return result, err
}

// EncounterConditionValue returns a single encounter condition value
//
//	(by name or ID).
func EncounterConditionValue(id string) (result models.EncounterConditionValue,
	err error) {
	err = request.Do(fmt.Sprintf("encounter-condition-value/%s", id), &result)
	return result, err
}
