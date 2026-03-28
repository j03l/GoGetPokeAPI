package encounters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncounterMethod(t *testing.T) {
	result, _ := EncounterMethod("1")
	assert.Equal(t, "walk", result.Name,
		"Expect to receive Walk.")
}

func TestEncounterMethodByName(t *testing.T) {
	result, _ := EncounterMethod("walk")
	assert.Equal(t, "walk", result.Name,
		"Expect to receive Walk.")
}

func TestEncounterMethodFail(t *testing.T) {
	result, _ := EncounterMethod("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestEncounterCondition(t *testing.T) {
	result, _ := EncounterCondition("1")
	assert.Equal(t, "swarm", result.Name,
		"Expect to receive Swarm.")
}

func TestEncounterConditionByName(t *testing.T) {
	result, _ := EncounterCondition("swarm")
	assert.Equal(t, "swarm", result.Name,
		"Expect to receive Swarm.")
}

func TestEncounterConditionFail(t *testing.T) {
	result, _ := EncounterCondition("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestEncounterConditionValue(t *testing.T) {
	result, _ := EncounterConditionValue("1")
	assert.Equal(t, "swarm-yes", result.Name,
		"Expect to receive Swarm (yes).")
}

func TestEncounterConditionValueByName(t *testing.T) {
	result, _ := EncounterConditionValue("swarm-yes")
	assert.Equal(t, "swarm-yes", result.Name,
		"Expect to receive Swarm (yes).")
}

func TestEncounterConditionValueFail(t *testing.T) {
	result, _ := EncounterConditionValue("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
