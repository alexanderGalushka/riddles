package routers

import (
	"fmt"
	consts "github.com/alexanderGalushka/riddles/api/constants"
	h "github.com/alexanderGalushka/riddles/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// InitRoutes initializes routes
func InitRoutes() *gin.Engine {
	uriPrefix := fmt.Sprintf("%s/%s", consts.Version, consts.ServiceName)

	r := gin.New()
	config := cors.DefaultConfig()
    config.AllowHeaders = []string{"Origin", "Access-Control-Allow-Origin"}
	config.AllowOrigins = []string{"http://localhost:63342"}
	config.ExposeHeaders = []string{"Content-Length", "Origin", "Access-Control-Allow-Origin"}

	r.Use(cors.New(config))
	r.GET(consts.HealthcheckURI, func(c *gin.Context) {
		h.GetHealthStatus(c)
	})

	api := r.Group(fmt.Sprintf("%s/:%s", uriPrefix, consts.RiddleTypeURIParam))
	api.GET("/", h.GetRiddleSolution)

	return r
}
