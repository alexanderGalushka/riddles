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

// TransferYtoX is the function to transfer contents of Y to X
func (jc *JugContainer) TransferYtoX() {
	waterAmountXCanTake := jc.jugX.totalVolume - jc.jugX.currentVolume
	// take the min value between the amount jug X can take and what jug Y currently has
	waterAmountToTransferFromYtoX := u.Min(waterAmountXCanTake, jc.jugY.currentVolume)
	jc.jugX.currentVolume += waterAmountToTransferFromYtoX
	jc.jugY.currentVolume -= waterAmountToTransferFromYtoX
}

// TransferXtoY is the function to transfer contents of X to Y
func (jc *JugContainer) TransferXtoY() {
	waterAmountYCanTake := jc.jugY.totalVolume - jc.jugY.currentVolume
	// take the min value between the amount jug Y can take and what jug X currently has
	waterAmountToTransferFromXtoY := u.Min(waterAmountYCanTake, jc.jugX.currentVolume)
	jc.jugY.currentVolume += waterAmountToTransferFromXtoY
	jc.jugX.currentVolume -= waterAmountToTransferFromXtoY
}

// AddStep is the function to add solution step
func (jc *JugContainer) AddStep(state string) {
	step := Step{consts.XKey: jc.jugX.currentVolume, consts.YKey: jc.jugY.currentVolume, consts.StateKey: state}
	jc.steps = append(jc.steps, step)
}
