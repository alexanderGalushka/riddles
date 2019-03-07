// +build unit

package handlers

import (
	consts "github.com/alexanderGalushka/riddles/api/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartWithBigJugToFindSolution(t *testing.T) {
	assert := assert.New(t)

	z := 4
	expectedStepsNum := 6

	totalVolumeX := 5
	currentVolumeX := 0

	totalVolumeY := 3
	currentVolumeY := 0

	jugX := Jug{
		totalVolume:   totalVolumeX,
		currentVolume: currentVolumeX,
	}

	jugY := Jug{
		totalVolume:   totalVolumeY,
		currentVolume: currentVolumeY,
	}

	jc := JugContainer{
		jugX: jugX,
		jugY: jugY,
	}

	t.Run("unit=StartWithBigJugToFindSolution", func(t *testing.T) {
		steps := startWithBigJugToFindSolution(jc, z)
		assert.Equal(expectedStepsNum, len(steps))
		assert.Equal(z, steps[len(steps)-1][consts.XKey].(int))
	})
}


func TestGetGreatestCommonDivisor(t *testing.T) {
	assert := assert.New(t)
	t.Run("unit=getGreatestCommonDivisor2", func(t *testing.T) {
		assert.Equal(2, getGreatestCommonDivisor(4, 2))
	})
	t.Run("unit=getGreatestCommonDivisor1", func(t *testing.T) {
		assert.Equal(1, getGreatestCommonDivisor(5, 3))
	})
}