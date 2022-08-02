package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store/services/domain/user/usecase"
	"store/services/models"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return UserHandler{userUsecase: userUsecase}
}

func (u *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userUsecase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "faild",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   user,
	})

}

func (u *UserHandler) AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "faild",
		})
		return
	}
	err := u.userUsecase.AddUser(user)
	if err != nil {
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "faild",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {

}

func (u *UserHandler) DeleteUser(c *gin.Context) {

}
