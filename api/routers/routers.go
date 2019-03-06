package routers

import (
	"fmt"
	consts "github.com/alexanderGalushka/riddles/api/constants"
	h "github.com/alexanderGalushka/riddles/api/handlers"
	"github.com/gin-gonic/gin"
)

// InitRoutes initializes routes
func InitRoutes() *gin.Engine {
	baseURI := fmt.Sprintf("%s/%s", consts.Version, consts.ServiceName)

	r := gin.New()
	r.Use(CORSMiddleware())

	r.GET(consts.HealthcheckURI, func(c *gin.Context) {
		h.GetHealthStatus(c)
	})

	riddleSolutionURI := fmt.Sprintf("%s/:%s/solution", baseURI, consts.RiddleTypeURIParam)

	r.GET(riddleSolutionURI, func(c *gin.Context) {
		h.GetRiddleSolution(c)
	})

	return r
}
