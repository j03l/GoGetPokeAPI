package evolution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvolutionChain(t *testing.T) {
	result, _ := EvolutionChain("1")
	assert.Equal(t, "bulbasaur", result.Chain.Species.Name,
		"Expect to receive Bulbasaur.")
}

func TestEvolutionChainFail(t *testing.T) {
	result, _ := EvolutionChain("asdf")
	assert.Equal(t, "", result.Chain.Species.Name,
		"Expect to receive an empty result.")
}

func TestEvolutionTrigger(t *testing.T) {
	result, _ := EvolutionTrigger("1")
	assert.Equal(t, "level-up", result.Name,
		"Expect to receive Level Up.")
}

func TestEvolutionTriggerByName(t *testing.T) {
	result, _ := EvolutionTrigger("level-up")
	assert.Equal(t, "level-up", result.Name,
		"Expect to receive Level Up.")
}

func TestEvolutionTriggerFail(t *testing.T) {
	result, _ := EvolutionTrigger("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
