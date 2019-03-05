package utils

import (
	"errors"
	consts "github.com/alexanderGalushka/riddles/api/constants"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetURIParam retrieves URI param from gin context
func GetURIParam(c *gin.Context, param string) (string, error) {
	p := c.Params.ByName(param)
	if p == consts.EmptyString {
		return consts.EmptyString, errors.New(param + " query parameter can not be empty")
	}
	return p, nil
}


// GetQueryIntParam retrieves query param from gin context and converts it to int
func GetQueryIntParam(c *gin.Context, param string) (int, error) {
	p := c.Query(param)
	if p == consts.EmptyString {
		return 0, errors.New(param + " query parameter can not be empty")
	}
	return strconv.Atoi(p)
}
