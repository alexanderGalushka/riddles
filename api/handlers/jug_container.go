package handlers

import (
	consts "github.com/alexanderGalushka/riddles/api/constants"
	u "github.com/alexanderGalushka/riddles/api/utils"
)

// Step is an alias for map[string]interface{}, representing each step of the riddle
type Step map[string]interface{}

// Jug is the struct for jug represented by its total capacity aka total volume and jug current state aka current volume
type Jug struct {
	totalVolume, currentVolume int
	id                         string
}

// JugContainer is the struct represented by two jugs, X and Y
type JugContainer struct {
	jugX, jugY Jug
	steps      []Step
}

// Empty is the function to zero out jug's current volume
func (j *Jug) Empty() {
	j.currentVolume = 0
}

// Fill is the function to fill up jug's current volume to its max capacity
func (j *Jug) Fill() {
	j.currentVolume = j.totalVolume
}

// TransferFromBigtoSmall is the function to transfer contents from big jug to small jug
func (jc *JugContainer) TransferFromBigtoSmall() {
	bigJug := jc.BigJug()
	smallJug := jc.SmallJug()

	waterAmountSmallCanTake := smallJug.totalVolume - smallJug.currentVolume
	// take the min value between the amount of small jug can take and what big jug currently has
	waterAmountToTransferFromBigToSmall := u.Min(waterAmountSmallCanTake, bigJug.currentVolume)
	smallJug.currentVolume += waterAmountToTransferFromBigToSmall
	bigJug.currentVolume -= waterAmountToTransferFromBigToSmall
}

// AddStep is the function to add solution step
func (jc *JugContainer) AddStep(state string) {
	step := Step{consts.XKey: jc.jugX.currentVolume, consts.YKey: jc.jugY.currentVolume, consts.StateKey: state}
	jc.steps = append(jc.steps, step)
}

// IsSolved is the function to identify if goal measurement of jug container has been achieved
func (jc *JugContainer) IsSolved(goalMeasurement int) bool {
	return jc.jugX.currentVolume == goalMeasurement ||
		jc.jugY.currentVolume == goalMeasurement ||
		jc.jugX.currentVolume+jc.jugY.currentVolume == goalMeasurement
}

// BigJug returns the greatest capacity jug of two
func (jc *JugContainer) BigJug() *Jug {
	if jc.jugX.totalVolume > jc.jugY.totalVolume {
		return &jc.jugX
	}
	if jc.jugY.totalVolume > jc.jugX.totalVolume {
		return &jc.jugY
	}
	// in case X equals Y, which is not intended by design
	return nil
}

// SmallJug returns the least capacity jug of two
func (jc *JugContainer) SmallJug() *Jug {
	if jc.jugX.totalVolume < jc.jugY.totalVolume {
		return &jc.jugX
	}
	if jc.jugY.totalVolume < jc.jugX.totalVolume {
		return &jc.jugY
	}
	// in case X equals Y, which is not intended by design
	return nil
}
