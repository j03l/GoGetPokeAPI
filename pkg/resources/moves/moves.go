package moves

import (
	"fmt"

	"github.com/j03l/GoGetPokeAPI/pkg/models"
)

// Move returns a single move (by name or ID).
func Move(id string) (result models.Move, err error) {
	err = do(fmt.Sprintf("move/%s", id), &result)
	return result, err
}

// MoveAilment returns a single move ailment (by name or ID).
func MoveAilment(id string) (result models.MoveAilment, err error) {
	err = do(fmt.Sprintf("move-ailment/%s", id), &result)
	return result, err
}

// MoveBattleStyle returns a single move battle style (by name or ID).
func MoveBattleStyle(id string) (result models.MoveBattleStyle, err error) {
	err = do(fmt.Sprintf("move-battle-style/%s", id), &result)
	return result, err
}

// MoveCategory returns a single move category (by name or ID).
func MoveCategory(id string) (result models.MoveCategory, err error) {
	err = do(fmt.Sprintf("move-category/%s", id), &result)
	return result, err
}

// MoveDamageClass returns a single move damage class (by name or ID).
func MoveDamageClass(id string) (result models.MoveDamageClass, err error) {
	err = do(fmt.Sprintf("move-damage-class/%s", id), &result)
	return result, err
}

// MoveLearnMethod returns a single move learn method (by name or ID).
func MoveLearnMethod(id string) (result models.MoveLearnMethod, err error) {
	err = do(fmt.Sprintf("move-learn-method/%s", id), &result)
	return result, err
}

// MoveTarget returns a single move target (by name or ID).
func MoveTarget(id string) (result models.MoveTarget, err error) {
	err = do(fmt.Sprintf("move-target/%s", id), &result)
	return result, err
}
