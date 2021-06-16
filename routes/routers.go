package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/go-admin/controllers"
	"github.com/radish-miyazaki/go-admin/middlewares"
)

func Setup(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)

		// ログイン済み（JWT認証済みだけアクセス可能）
		v1.Use(middlewares.Authenticate())
		{
			// Operation for Login User
			v1.GET("/user", controllers.User)
			v1.PUT("/user/info", controllers.UpdateInfo)
			v1.PUT("/user/password", controllers.UpdatePassword)
			v1.POST("/logout", controllers.Logout)

			// Get to all Permission data
			v1.GET("/permissions", controllers.AllPermissions)

			// Image Upload
			v1.POST("/upload", controllers.Upload)
			v1.Static("/uploads/", "./uploads")

			// Sales to date
			v1.GET("/chart", controllers.Chart)

			v1.Use(middlewares.Authorized()) // Permission Middleware
			{
				// CRUD for User models
				v1.GET("/users", controllers.AllUsers)
				v1.POST("/users", controllers.CreateUser)
				v1.GET("/users/:id", controllers.GetUser)
				v1.PUT("/users/:id", controllers.UpdateUser)
				v1.DELETE("/users/:id", controllers.DeleteUser)

				// CRUD for Role models
				v1.GET("/roles", controllers.AllRoles)
				v1.POST("/roles", controllers.CreateRole)
				v1.GET("/roles/:id", controllers.GetRole)
				v1.PUT("/roles/:id", controllers.UpdateRole)
				v1.DELETE("/roles/:id", controllers.DeleteRole)

				// CRUD for Product models
				v1.GET("/products", controllers.AllProducts)
				v1.POST("/products", controllers.CreateProduct)
				v1.GET("/products/:id", controllers.GetProduct)
				v1.PUT("/products/:id", controllers.UpdateProduct)
				v1.DELETE("/products/:id", controllers.DeleteProduct)

				// CRUD for Order models
				v1.GET("/orders", controllers.AllOrders)
				v1.POST("/orders/export", controllers.Export)
			}
		}
	}
}
