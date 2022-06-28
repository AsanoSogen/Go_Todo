package server

import (
	hello "github.com/Go_project/controller/hello"
	// user "github.com/Go_project/controller/user"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	// u := r.Group("/users")
	// {
	// 	ctrl := user.Controller{}
	// 	u.GET("", ctrl.Index)
	// 	u.GET("/:id", ctrl.Show)
	// 	u.POST("", ctrl.Create)
	// 	u.PUT("/:id", ctrl.Update)
	// 	u.DELETE("/:id", ctrl.Delete)
	// }

	h := r.Group("/hello")
	{
		ctrl := hello.Controller{}
		h.GET("", ctrl.HelloIndex)
	}
	return r
}
