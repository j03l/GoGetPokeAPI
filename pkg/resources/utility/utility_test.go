package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage(t *testing.T) {
	result, _ := Language("1")
	assert.Equal(t, "ja-Hrkt", result.Name,
		"Expect to receive Japanese.")
}

func TestLanguageByName(t *testing.T) {
	result, _ := Language("ja-Hrkt")
	assert.Equal(t, "ja-Hrkt", result.Name,
		"Expect to receive Japanese.")
}

func TestLanguageFail(t *testing.T) {
	result, _ := Language("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
