package http

import "github.com/gin-gonic/gin"

type UserRouter struct {
	handler UserHandler
}

func NewUserRouter(handler UserHandler) UserRouter {
	return UserRouter{handler}
}

func (u *UserRouter) UseRoute(r *gin.RouterGroup) {
	group := r.Group("users")
	group.GET("/:id", u.handler.GetUser)
	group.POST("/", u.handler.AddUser)
}
