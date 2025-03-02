package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zkfmapf123/go-llm/internal/utils"
)

const (
	prevWord = "장례식"
	useWord  = "식사"
)

func Test_userInputComparison(t *testing.T) {
	b := utils.ComparisonVarsPrefixSuffix("장례식", "식사")
	assert.Equal(t, b, false)

	c := utils.ComparisonVarsPrefixSuffix("장례식", "식장")
	assert.Equal(t, c, true)
}

func Test_FirstVoca(t *testing.T) {
	r := []rune(prevWord)

	first := string(r[0])
	se := string(r[len(r)-2])
	th := string(r[len(r)-1])

	assert.Equal(t, first, "장")
	assert.Equal(t, se, "례")
	assert.Equal(t, th, "식")

}
