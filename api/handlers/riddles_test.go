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

func TestSolveWaterJugRiddle(t *testing.T) {
	assert := assert.New(t)
	t.Run("unit=solveWaterJugRiddleX=3_Y=4_Z=3", func(t *testing.T) {
		z := 3
		steps, err := solveWaterJugRiddle(3, 4, z)
		assert.Nil(err)
		assert.Equal(1, len(steps))
		s := steps[consts.StepsKey].([]Step)
		assert.Equal(z, s[0][consts.XKey].(int))
		assert.Equal(0, s[0][consts.YKey].(int))
	})
	t.Run("unit=solveWaterJugRiddleX=4_Y=3_Z=3", func(t *testing.T) {
		z := 3
		steps, err := solveWaterJugRiddle(4, 3, z)
		assert.Nil(err)
		assert.Equal(1, len(steps))
		s := steps[consts.StepsKey].([]Step)
		assert.Equal(z, s[0][consts.YKey].(int))
		assert.Equal(0, s[0][consts.XKey].(int))
	})
	t.Run("unit=solveWaterJugRiddleX=2_Y=1_Z=3", func(t *testing.T) {
		z := 3
		steps, err := solveWaterJugRiddle(2, 1, z)
		assert.Nil(err)
		assert.Equal(1, len(steps))
		s := steps[consts.StepsKey].([]Step)
		assert.Equal(z, s[0][consts.YKey].(int)+s[0][consts.XKey].(int))
	})
	t.Run("unit=solveWaterJugRiddleX=2_Y=1_Z=0", func(t *testing.T) {
		z := 0
		steps, err := solveWaterJugRiddle(2, 1, z)
		assert.Nil(err)
		assert.Equal(1, len(steps))
		s := steps[consts.StepsKey].([]Step)
		assert.Equal(0, s[0][consts.XKey].(int))
		assert.Equal(0, s[0][consts.YKey].(int))
	})
	t.Run("unit=solveWaterJugRiddleX=2_Y=1_Z=0", func(t *testing.T) {
		z := 4
		steps, err := solveWaterJugRiddle(2, 1, z)
		assert.Nil(err)
		assert.Equal(1, len(steps))
		s := steps[consts.StepsKey].([]Step)
		assert.Equal(0, s[0][consts.XKey].(int))
		assert.Equal(0, s[0][consts.YKey].(int))
	})
}
