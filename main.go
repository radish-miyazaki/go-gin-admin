package main

import (
	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/go-admin/cors"
	"github.com/radish-miyazaki/go-admin/db"
	"github.com/radish-miyazaki/go-admin/routes"
)

func main() {
	r := gin.Default() // Engine, Middleware
	routes.Setup(r)    // Router
	db.Connect()       // DB
	cors.Setup(r)      // CORS

	if err := r.Run(":8080"); err != nil {
		panic("couldn't start api server!")
	}
}
