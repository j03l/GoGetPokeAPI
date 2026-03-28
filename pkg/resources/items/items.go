package items

import (
	"fmt"

	"github.com/j03l/GoGetPokeAPI/internal/request"
	"github.com/j03l/GoGetPokeAPI/pkg/models"
)

// Item returns a single item (by name or ID).
func Item(id string) (result models.Item, err error) {
	err = request.Do(fmt.Sprintf("item/%s", id), &result)
	return result, err
}

// ItemAttribute returns a single item attribute (by name or ID).
func ItemAttribute(id string) (result models.ItemAttribute, err error) {
	err = request.Do(fmt.Sprintf("item-attribute/%s", id), &result)
	return result, err
}

// ItemCategory returns a single item category (by name or ID).
func ItemCategory(id string) (result models.ItemCategory, err error) {
	err = request.Do(fmt.Sprintf("item-category/%s", id), &result)
	return result, err
}

// ItemFlingEffect returns a single item fling effect (by name or ID).
func ItemFlingEffect(id string) (result models.ItemFlingEffect, err error) {
	err = request.Do(fmt.Sprintf("item-fling-effect/%s", id), &result)
	return result, err
}

// ItemPocket returns a single item pocket (by name or ID).
func ItemPocket(id string) (result models.ItemPocket, err error) {
	err = request.Do(fmt.Sprintf("item-pocket/%s", id), &result)
	return result, err
}
