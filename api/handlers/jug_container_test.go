// +build unit

package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	assert := assert.New(t)

	totalVolume := 5
	currentVolume := 3

	jugX := Jug{
		totalVolume:   totalVolume,
		currentVolume: currentVolume,
	}

	t.Run("unit=EmptyIndependentJug", func(t *testing.T) {

		jugX.Empty()
		assert.Equal(0, jugX.currentVolume)
		assert.Equal(totalVolume, jugX.totalVolume)
	})
	t.Run("unit=EmptyJugInTheContainer", func(t *testing.T) {
		jc := JugContainer{
			jugX: jugX,
		}
		jc.jugX.Empty()
		assert.Equal(0, jc.jugX.currentVolume)
		assert.Equal(totalVolume, jc.jugX.totalVolume)
	})
}

func TestFill(t *testing.T) {
	assert := assert.New(t)

	totalVolume := 5
	currentVolume := 0

	jugX := Jug{
		totalVolume:   totalVolume,
		currentVolume: currentVolume,
	}

	t.Run("unit=FillIndependentJug", func(t *testing.T) {

		jugX.Fill()
		assert.Equal(totalVolume, jugX.currentVolume)
	})
	t.Run("unit=FillJugInTheContainer", func(t *testing.T) {
		jc := JugContainer{
			jugX: jugX,
		}
		jc.jugX.Fill()
		assert.Equal(totalVolume, jc.jugX.currentVolume)
	})
}

func TestTransferFromBigtoSmall(t *testing.T) {
	assert := assert.New(t)
	totalVolumeX := 5
	currentVolumeX := 2

	totalVolumeY := 3
	currentVolumeY := 1

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

	t.Run("unit=TransferFromBigToSmallWithinContainer", func(t *testing.T) {
		jc.TransferFromBigtoSmall()
		assert.Equal(currentVolumeX+currentVolumeY, jc.jugY.currentVolume)
		assert.Equal(0, jc.jugX.currentVolume)
	})
}

func TestIsSolved(t *testing.T) {
	assert := assert.New(t)

	totalVolumeX := 5
	currentVolumeX := 2

	totalVolumeY := 3
	currentVolumeY := 1

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

	t.Run("unit=IsSolvedFalse", func(t *testing.T) {
		assert.False(jc.IsSolved(4))
	})
	t.Run("unit=IsSolvedTrueY", func(t *testing.T) {
		assert.True(jc.IsSolved(currentVolumeY))
	})
	t.Run("unit=IsSolvedTrueX", func(t *testing.T) {
		assert.True(jc.IsSolved(currentVolumeX))
	})
	t.Run("unit=IsSolvedTrueBoth", func(t *testing.T) {
		assert.True(jc.IsSolved(currentVolumeX+currentVolumeY))
	})
}