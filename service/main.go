package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"ytool/router"
)

func main() {
	r := gin.Default()
	r.Use(Cors())
	r.Use(static.Serve("/", static.LocalFile("./webapp", true)))
	router.InitRouter(r)

	log.Println("start server...")
	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Origin") != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
