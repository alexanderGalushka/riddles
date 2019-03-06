package handlers

import u "github.com/alexanderGalushka/riddles/api/utils"

// Step is an alias for map[string]interface{}, representing each step of the riddle
type Step map[string]interface{}

// Jug is the struct for jug represented by its total capacity aka total volume and jug current state aka current volume
type Jug struct {
	totalVolume, currentVolume int
}

// JugContainer is the struct represented by two jugs, X and Y
type JugContainer struct {
	jugX, jugY Jug
	steps []Step
}

func (j *Jug) Empty() {
	j.currentVolume = 0
}

func (j *Jug) Fill() {
	j.currentVolume = j.totalVolume
}

func (jc *JugContainer) TransferYtoX() {
	waterAmountXCanTake := jc.jugX.totalVolume - jc.jugX.currentVolume
	// take the min value between the amount jug X can take and what jug Y currently has
	waterAmountToTransferFromYtoX := u.Min(waterAmountXCanTake, jc.jugY.currentVolume)
	jc.jugX.currentVolume += waterAmountToTransferFromYtoX
	jc.jugY.currentVolume -= waterAmountToTransferFromYtoX
}

func (jc *JugContainer) TransferXtoY() {
	waterAmountYCanTake := jc.jugY.totalVolume - jc.jugY.currentVolume
	// take the min value between the amount jug Y can take and what jug X currently has
	waterAmountToTransferFromXtoY := u.Min(waterAmountYCanTake, jc.jugX.currentVolume)
	jc.jugY.currentVolume += waterAmountToTransferFromXtoY
	jc.jugX.currentVolume -= waterAmountToTransferFromXtoY
}

func (jc *JugContainer) AddStep(state string) {
	step := Step{"x": jc.jugX.currentVolume, "y": jc.jugY.currentVolume, "state": state}
	jc.steps = append(jc.steps, step)
}
