package conf

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

var Ginner *gin.Engine

func init() {
	Ginner = gin.Default()
	Ginner.Use(Cors())
	Ginner.Use(gin.Recovery())
	Ginner.Use(static.Serve("/ui", static.LocalFile("./ui/build", true)))
}

var db = make(map[string]string)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
				log.Printf("Panic info is: %s", debug.Stack())
			}
		}()
		c.Next()
	}
}
