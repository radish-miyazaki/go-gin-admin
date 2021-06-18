package main

import (
	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/go-admin/cors"
	"github.com/radish-miyazaki/go-admin/db"
	"github.com/radish-miyazaki/go-admin/routes"
)

func main() {
	r := gin.Default() // Engine, Middleware
	/*
		 INFO: cors_middlewareはルーティング全体に適用する必要があるので、
				ルーティング設定の前に持ってくる必要がある。
	*/
	cors.Setup(r)
	db.Connect()    // DB
	routes.Setup(r) // Router

	if err := r.Run(":8080"); err != nil {
		panic("couldn't start api server!")
	}
}
