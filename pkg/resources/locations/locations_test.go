package locations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	result, _ := Location("1")
	assert.Equal(t, "canalave-city", result.Name,
		"Expect to receive Canalave City.")
}

func TestLocationByName(t *testing.T) {
	result, _ := Location("canalave-city")
	assert.Equal(t, "canalave-city", result.Name,
		"Expect to receive Canalave City.")
}

func TestLocationFail(t *testing.T) {
	result, _ := Location("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestLocationArea(t *testing.T) {
	result, _ := LocationArea("1")
	assert.Equal(t, "canalave-city-area", result.Name,
		"Expect to receive Canalave City area.")
}

func TestLocationAreaByName(t *testing.T) {
	result, _ := LocationArea("canalave-city-area")
	assert.Equal(t, "canalave-city-area", result.Name,
		"Expect to receive Canalave City area.")
}

func TestLocationAreaFail(t *testing.T) {
	result, _ := LocationArea("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPalParkArea(t *testing.T) {
	result, _ := PalParkArea("1")
	assert.Equal(t, "forest", result.Name,
		"Expect to receive Forest.")
}

func TestPalParkAreaByName(t *testing.T) {
	result, _ := PalParkArea("forest")
	assert.Equal(t, "forest", result.Name,
		"Expect to receive Forest.")
}

func TestPalParkAreaFail(t *testing.T) {
	result, _ := PalParkArea("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestRegion(t *testing.T) {
	result, _ := Region("1")
	assert.Equal(t, "kanto", result.Name,
		"Expect to receive Kanto.")
}

func TestRegionByName(t *testing.T) {
	result, _ := Region("kanto")
	assert.Equal(t, "kanto", result.Name,
		"Expect to receive Kanto.")
}

func TestRegionFail(t *testing.T) {
	result, _ := Region("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
