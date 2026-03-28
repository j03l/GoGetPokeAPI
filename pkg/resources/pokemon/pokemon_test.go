package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbility(t *testing.T) {
	result, _ := Ability("1")
	assert.Equal(t, "stench", result.Name,
		"Expect to receive Stench.")
}

func TestAbilityByName(t *testing.T) {
	result, _ := Ability("stench")
	assert.Equal(t, "stench", result.Name,
		"Expect to receive Stench.")
}

func TestAbilityFail(t *testing.T) {
	result, _ := Ability("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestCharacteristic(t *testing.T) {
	result, _ := Characteristic("1")
	for _, description := range result.Descriptions {
		if description.Language.Name == "en" {
			assert.Equal(t, "Loves to eat", description.Description,
				"Expect to receive a description.")
		}
	}
}

func TestCharacteristicFail(t *testing.T) {
	result, _ := Characteristic("asdf")
	assert.Equal(t, 0, len(result.Descriptions),
		"Expect to receive an empty result.")
}

func TestEggGroup(t *testing.T) {
	result, _ := EggGroup("1")
	assert.Equal(t, "monster", result.Name,
		"Expect to receive Monster.")
}

func TestEggGroupByName(t *testing.T) {
	result, _ := EggGroup("monster")
	assert.Equal(t, "monster", result.Name,
		"Expect to receive Monster.")
}

func TestEggGroupFail(t *testing.T) {
	result, _ := EggGroup("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestGender(t *testing.T) {
	result, _ := Gender("1")
	assert.Equal(t, "female", result.Name,
		"Expect to receive Female.")
}

func TestGenderByName(t *testing.T) {
	result, _ := Gender("female")
	assert.Equal(t, "female", result.Name,
		"Expect to receive Female.")
}

func TestGenderFail(t *testing.T) {
	result, _ := Gender("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestGrowthRate(t *testing.T) {
	result, _ := GrowthRate("1")
	assert.Equal(t, "slow", result.Name,
		"Expect to receive Slow.")
}

func TestGrowthRateByName(t *testing.T) {
	result, _ := GrowthRate("slow")
	assert.Equal(t, "slow", result.Name,
		"Expect to receive Slow.")
}

func TestGrowthRateFail(t *testing.T) {
	result, _ := GrowthRate("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestNature(t *testing.T) {
	result, _ := Nature("1")
	assert.Equal(t, "hardy", result.Name,
		"Expect to receive Hardy.")
}

func TestNatureByName(t *testing.T) {
	result, _ := Nature("hardy")
	assert.Equal(t, "hardy", result.Name,
		"Expect to receive Hardy.")
}

func TestNatureFail(t *testing.T) {
	result, _ := Nature("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokeathlonStat(t *testing.T) {
	result, _ := PokeathlonStat("1")
	assert.Equal(t, "speed", result.Name,
		"Expect to receive Speed.")
}

func TestPokeathlonStatByName(t *testing.T) {
	result, _ := PokeathlonStat("speed")
	assert.Equal(t, "speed", result.Name,
		"Expect to receive Speed.")
}

func TestPokeathlonStatFail(t *testing.T) {
	result, _ := PokeathlonStat("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemon(t *testing.T) {
	result, _ := Pokemon("1")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonByName(t *testing.T) {
	result, _ := Pokemon("bulbasaur")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFail(t *testing.T) {
	result, _ := Pokemon("digimon")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonColor(t *testing.T) {
	result, _ := PokemonColor("1")
	assert.Equal(t, "black", result.Name,
		"Expect to receive Black.")
}

func TestPokemonColorByName(t *testing.T) {
	result, _ := PokemonColor("black")
	assert.Equal(t, "black", result.Name,
		"Expect to receive Black.")
}

func TestPokemonColorFail(t *testing.T) {
	result, _ := PokemonColor("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonForm(t *testing.T) {
	result, _ := PokemonForm("1")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFormByName(t *testing.T) {
	result, _ := PokemonForm("bulbasaur")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFormFail(t *testing.T) {
	result, _ := PokemonForm("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonHabitat(t *testing.T) {
	result, _ := PokemonHabitat("1")
	assert.Equal(t, "cave", result.Name,
		"Expect to receive Cave.")
}

func TestPokemonHabitatByName(t *testing.T) {
	result, _ := PokemonHabitat("cave")
	assert.Equal(t, "cave", result.Name,
		"Expect to receive Cave.")
}

func TestPokemonHabitatFail(t *testing.T) {
	result, _ := PokemonHabitat("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonShape(t *testing.T) {
	result, _ := PokemonShape("1")
	assert.Equal(t, "ball", result.Name,
		"Expect to receive Ball.")
}

func TestPokemonShapeByName(t *testing.T) {
	result, _ := PokemonShape("ball")
	assert.Equal(t, "ball", result.Name,
		"Expect to receive Ball.")
}

func TestPokemonShapeFail(t *testing.T) {
	result, _ := PokemonShape("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonSpecies(t *testing.T) {
	result, _ := PokemonSpecies("1")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonSpeciesByName(t *testing.T) {
	result, _ := PokemonSpecies("bulbasaur")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonSpeciesFail(t *testing.T) {
	result, _ := PokemonSpecies("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestStat(t *testing.T) {
	result, _ := Stat("1")
	assert.Equal(t, "hp", result.Name,
		"Expect to receive HP.")
}

func TestStatByName(t *testing.T) {
	result, _ := Stat("hp")
	assert.Equal(t, "hp", result.Name,
		"Expect to receive HP.")
}

func TestStatFail(t *testing.T) {
	result, _ := Stat("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestType(t *testing.T) {
	result, _ := Type("1")
	assert.Equal(t, "normal", result.Name,
		"Expect to receive Normal.")
}

func TestTypeByName(t *testing.T) {
	result, _ := Type("normal")
	assert.Equal(t, "normal", result.Name,
		"Expect to receive Normal.")
}

func TestTypeFail(t *testing.T) {
	result, _ := Type("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
