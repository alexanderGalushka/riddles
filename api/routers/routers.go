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
	r.Use(CORSMiddleware())

	r.GET(consts.HealthcheckURI, func(c *gin.Context) {
		h.GetHealthStatus(c)
	})

	baseURI := fmt.Sprintf("%s/:%s", uriPrefix, consts.RiddleTypeURIParam)

	r.GET(baseURI, func(c *gin.Context) {
		h.GetRiddleSolution(c)
	})

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, Cache-Control")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
