package handlers

import (
	"fmt"
	consts "github.com/alexanderGalushka/riddles/api/constants"
	u "github.com/alexanderGalushka/riddles/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const noSolutionErrorTemplate = "no solution, unable to measure %d gallons of water with jug X capacity of %d and jug Y capacity of %d"

// GetRiddleSolution returns step by step solution for supported riddle problem
func GetRiddleSolution(c *gin.Context) {

	riddleType, err := u.GetURIParam(c, consts.RiddleTypeURIParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{consts.ErrorKey: err.Error()})
	}

	result, err := solveRiddle(riddleType, c)
	if err != nil {
		// TODO
		c.JSON(http.StatusBadRequest, gin.H{consts.ErrorKey: err.Error()})
	}
	c.JSON(http.StatusOK, result)

	//if riddleTypeFlag, ok := supportedRiddleType[riddleType]; ok {
	//	if !riddleTypeFlag {
	//		c.JSON(http.StatusNotImplemented,
	//			gin.H{consts.ErrorKey: fmt.Sprintf("handler for %s riddle type has not been implemented yet",
	//				riddleType)})
	//	} else {
	//		result, err := solveRiddle(riddleType, c)
	//		if err != nil {
	//			c.JSON(http.StatusInternalServerError,
	//				gin.H{consts.ErrorKey: err.Error()})
	//		}
	//		c.JSON(http.StatusOK, result)
	//	}
	//} else {
	//	c.JSON(http.StatusUnprocessableEntity,
	//		gin.H{consts.ErrorKey: fmt.Sprintf("%s is unknown riddle type", riddleType)})
	//}

}

func solveRiddle(riddleType string, c *gin.Context) (map[string]interface{}, error) {
	switch riddleType {
	case "water_jug":
		x, err := u.GetQueryIntParam(c, "x")
		if err != nil {
			return nil, err
		}
		y, err := u.GetQueryIntParam(c, "y")
		if err != nil {
			return nil, err
		}
		z, err := u.GetQueryIntParam(c, "z")
		if err != nil {
			return nil, err
		}
		return solveWaterJugRiddle(x, y, z)
	case "egg_equation":
		return nil, fmt.Errorf("handler for %s riddle type has not been implemtned yet", riddleType)
	default:
		return nil, fmt.Errorf("%s is unknown riddle type", riddleType)
	}

}

func solveWaterJugRiddle(inputX int, inputY int, inputZ int) (map[string]interface{}, error) {
	xKey := "x"
	yKey := "y"
	if inputZ == 0 {
		return singleStepTwoParamWaterJugSolution(xKey, inputX, yKey, inputY, "No need to solve, goal is set to 0"), nil
	}
	if inputZ == inputX {
		return singleStepWaterJugSolution(xKey, inputX, fmt.Sprintf("Fill up X with %d gallons of water", inputX)), nil
	}
	if inputZ == inputY {
		return singleStepWaterJugSolution(yKey, inputY, fmt.Sprintf("Fill up Y with %d gallons of water", inputY)), nil
	}
	if inputZ == inputX+inputY {
		return singleStepTwoParamWaterJugSolution(xKey, inputX, yKey, inputY, fmt.Sprintf("Fill up X with %d and Y with %d gallons of water", inputX, inputY)), nil
	}
	if inputX+inputY < inputZ {
		return nil, fmt.Errorf(noSolutionErrorTemplate, inputZ, inputX, inputZ)
	}
	if !validateWithBezoutsIdentity(inputX, inputY, inputZ) {
		return nil, fmt.Errorf(noSolutionErrorTemplate, inputZ, inputX, inputZ)
	}
	jugX := Jug{
		totalVolume:   inputX,
		currentVolume: 0,
	}
	jugY := Jug{
		totalVolume:   inputY,
		currentVolume: 0,
	}
	jc := JugContainer{
		jugX: jugX,
		jugY: jugY,
		steps: []Step{},
	}
	steps := findAllSteps(jc, inputZ)
	result := map[string]interface{}{"steps": steps}
	return result, nil
}


func findAllSteps(jc JugContainer, z int) []Step {

	// very first step: keep the smaller jug empty and fill up the larger one
	if jc.jugY.totalVolume > jc.jugX.totalVolume {
		jc.jugY.Fill()
		jc.AddStep("Fill up Y")
		for {
			jc.TransferYtoX()
			jc.AddStep("Transfer from Y to fill up X")
			if jc.jugY.currentVolume == z || jc.jugX.currentVolume == z || jc.jugY.currentVolume+jc.jugX.currentVolume == z {
				return jc.steps
			}
			jc.jugX.Empty()
			jc.AddStep("Empty X")
			if jc.jugY.currentVolume < jc.jugX.totalVolume {
				jc.TransferYtoX()
				jc.AddStep("Transfer from Y to X")
				jc.jugY.Fill()
				jc.AddStep( "Fill up from Y")
			}
		}
	}
	if jc.jugX.totalVolume > jc.jugY.totalVolume {
		jc.jugX.Fill()
		jc.AddStep("Fill up X")
		for {
			jc.TransferXtoY()
			jc.AddStep("Transfer from X to fill up Y")
			if jc.jugX.currentVolume == z || jc.jugY.currentVolume == z || jc.jugX.currentVolume+jc.jugY.currentVolume == z {
				return jc.steps
			}
			jc.jugY.Empty()
			jc.AddStep("Empty Y")
			if jc.jugX.currentVolume < jc.jugY.totalVolume {
				jc.TransferXtoY()
				jc.AddStep("Transfer from X to Y")
				jc.jugX.Fill()
				jc.AddStep( "Fill up from X")
			}
		}
	}
	return jc.steps
}

// see https://en.wikipedia.org/wiki/Bézout%27s_identity
func validateWithBezoutsIdentity(inputX int, inputY int, inputZ int) bool {
	gcd := getGreatestCommonDivisor(inputX, inputY)
	return inputZ%gcd == 0
}

// see https://oeis.org/wiki/Euclidean_algorithm
func getGreatestCommonDivisor(x int, y int) int {
	for y != 0 {
		temp := y
		y = x % y
		x = temp
	}
	return x
}

func singleStepWaterJugSolution(idKey string, value int, state string) map[string]interface{} {
	var steps []Step
	steps = append(steps, Step{idKey: value, consts.StateKey: state})
	return presentFinalResult(steps)
}

func singleStepTwoParamWaterJugSolution(idKey1 string, value1 int, idKey2 string, value2 int, state string) map[string]interface{} {
	var steps []Step
	steps = append(steps, Step{idKey1: value1, idKey2: value2, consts.StateKey: state})
	return presentFinalResult(steps)
}

func presentFinalResult(steps interface{}) map[string]interface{} {
	return map[string]interface{}{"steps": steps}
}
