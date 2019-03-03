package routers

import (
	"fmt"
	consts "github.com/alexanderGalushka/riddles/api/constants"
	h "github.com/alexanderGalushka/riddles/api/handlers"
	"github.com/gin-gonic/gin"
)

// InitRoutes initializes routes
func InitRoutes() *gin.Engine {
	uriPrefix := fmt.Sprintf("%s/%s", consts.Version, consts.ServiceName)

	r := gin.New()
	r.GET(consts.HealthcheckURI, func(c *gin.Context) {
		h.GetHealthStatus(c)
	})

	api := r.Group(fmt.Sprintf("%s/:%s", uriPrefix, consts.RiddleTypeURIParam))
	api.GET("/", h.GetRiddleSolution)

	return r
}
