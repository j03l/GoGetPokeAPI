package berries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBerry(t *testing.T) {
	result, _ := Berry("1")
	assert.Equal(t, "cheri", result.Name,
		"Expect to receive Cheri.")
}

func TestBerryByName(t *testing.T) {
	result, _ := Berry("cheri")
	assert.Equal(t, "cheri", result.Name,
		"Expect to receive Cheri.")
}

func TestBerryFail(t *testing.T) {
	result, _ := Berry("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestBerryFirmness(t *testing.T) {
	result, _ := BerryFirmness("1")
	assert.Equal(t, "very-soft", result.Name,
		"Expect to receive Very Soft.")
}

func TestBerryFirmnessByName(t *testing.T) {
	result, _ := BerryFirmness("very-soft")
	assert.Equal(t, "very-soft", result.Name,
		"Expect to receive Very Soft.")
}

func TestBerryFirmnessFail(t *testing.T) {
	result, _ := BerryFirmness("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestBerryFlavor(t *testing.T) {
	result, _ := BerryFlavor("1")
	assert.Equal(t, "spicy", result.Name,
		"Expect to receive Spicy.")
}

func TestBerryFlavorByName(t *testing.T) {
	result, _ := BerryFlavor("spicy")
	assert.Equal(t, "spicy", result.Name,
		"Expect to receive Spicy.")
}

func TestBerryFlavorFail(t *testing.T) {
	result, _ := BerryFlavor("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
