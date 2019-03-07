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
		x, err := u.GetQueryIntParam(c, consts.XKey)
		if err != nil {
			return nil, err
		}
		y, err := u.GetQueryIntParam(c, consts.YKey)
		if err != nil {
			return nil, err
		}
		z, err := u.GetQueryIntParam(c, "z")
		if err != nil {
			return nil, err
		}
		return solveWaterJugRiddle(x, y, z)
	case "egg_equation":
		return nil, fmt.Errorf("handler for %s riddle type has not been implemented", riddleType)
	default:
		return nil, fmt.Errorf("%s is unknown riddle type", riddleType)
	}

}

func solveWaterJugRiddle(inputX int, inputY int, inputZ int) (map[string]interface{}, error) {
	// simple cases
	switch inputZ {
	case 0:
		return singleStepTwoParamWaterJugSolution(consts.XKey, inputX, consts.YKey, inputY, "No need to solve, goal is set to 0"), nil
	case inputX:
		return singleStepTwoParamWaterJugSolution(consts.XKey, inputX, consts.YKey, inputY, "No need to solve, goal is set to 0"), nil
	case inputY:
		return singleStepWaterJugSolution(consts.YKey, inputY, fmt.Sprintf("Fill up Y with %d gallons of water", inputY)), nil
	case inputX+inputY:
		return singleStepTwoParamWaterJugSolution(consts.XKey, inputX, consts.YKey, inputY, fmt.Sprintf("Fill up X with %d and Y with %d gallons of water", inputX, inputY)), nil
	}

	// no solution, validate if there is a solution before trying to compute steps
	if inputX+inputY < inputZ || !validateWithBezoutsIdentity(inputX, inputY, inputZ){
		return singleStepTwoParamWaterJugSolution(consts.XKey, inputX, consts.YKey, inputY, fmt.Sprintf(noSolutionErrorTemplate, inputZ, inputX, inputZ)), nil
	}

	jugX := Jug{
		totalVolume:   inputX,
		currentVolume: 0,
		id: "X",
	}
	jugY := Jug{
		totalVolume:   inputY,
		currentVolume: 0,
		id: "Y",
	}
	jc := JugContainer{
		jugX:  jugX,
		jugY:  jugY,
		steps: []Step{},
	}

	// compute steps
	steps := startWithBigJugToFindSolution(jc, inputZ)
	result := map[string]interface{}{consts.StepsKey: steps}
	return result, nil
}

func startWithBigJugToFindSolution(jc JugContainer, z int) []Step {
	bigJug := jc.BigJug()
	smallJug := jc.SmallJug()
	bigJug.Fill()
	jc.AddStep(fmt.Sprintf("Fill up %s", bigJug.id))
	for {
		jc.TransferFromBigtoSmall()
		jc.AddStep(fmt.Sprintf("Transfer from %s to fill up %s", bigJug.id, smallJug.id))
		if jc.IsSolved(z) {
			return jc.steps
		}
		smallJug.Empty()
		jc.AddStep(fmt.Sprintf("Empty %s", smallJug.id))
		if bigJug.currentVolume < smallJug.totalVolume {
			jc.TransferFromBigtoSmall()
			jc.AddStep(fmt.Sprintf("Transfer from %s to %s", bigJug.id, smallJug.id))
			bigJug.Fill()
			jc.AddStep(fmt.Sprintf("Fill up from %s", bigJug.id))
		}
	}
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
	return map[string]interface{}{consts.StepsKey: steps}
}
