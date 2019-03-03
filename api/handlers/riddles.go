package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	consts "github.com/alexanderGalushka/riddles/api/constants"
	u "github.com/alexanderGalushka/riddles/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Step is an alias for map[string]interface{}, representing each step of the riddle
type Step map[string]interface{}


var supportedRiddleType = map[string]bool{
	"water_jug":    true,
	"egg_equation": false,
}

// GetRiddleSolution returns step by step solution for supported riddle problem
func GetRiddleSolution(c *gin.Context) {
	riddleType, err := u.GetURIParam(c, consts.RiddleTypeURIParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{consts.ErrorKey: err.Error()})
	}
	if riddleTypeFlag, ok := supportedRiddleType[riddleType]; ok {
		if !riddleTypeFlag {
			c.JSON(http.StatusNotImplemented,
				gin.H{consts.ErrorKey: fmt.Sprintf("handler for %s riddle type has not been implemtned yet",
					riddleType)})
		} else {
			result, err := solveRiddle(riddleType)
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					gin.H{consts.ErrorKey: err.Error()})
			}
			c.JSON(http.StatusOK, result)
		}
	} else {
		c.JSON(http.StatusUnprocessableEntity,
			gin.H{consts.ErrorKey: fmt.Sprintf("%s is unknown riddle type", riddleType)})
	}

}

func solveRiddle(riddleType string) (string, error) {

}

func solveWaterJugRiddle(inputX int, inputY int, inputZ int) (string, error) {
	var steps []Step
    xKey := "x"
	yKey := "y"
	if inputZ == inputX {
		return singleStepWaterJugSolution(xKey, inputX)
	}
	if inputZ == inputY {
		return singleStepWaterJugSolution(yKey, inputY)
	}

	if inputZ > inputX && inputZ > inputY {
       return consts.EmptyString, errors.New("")
	} else {

	}
    if
	result := map[string]interface{}{"steps": steps}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return consts.EmptyString, err
	}
	return string(jsonResult), nil
}

func singleStepWaterJugSolution(idKey string, value int) (string, error) {
	var steps []Step
	steps = append(steps, Step{idKey: value, consts.StateKey: "filled"})
	result, err := presentFinalResult(steps)
	if err != nil {
		return consts.EmptyString, nil
	}
	return result, nil
}

func presentFinalResult(steps interface{}) (string, error) {
	result := map[string]interface{}{"steps": steps}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return consts.EmptyString, err
	}
	return string(jsonResult), nil
}
