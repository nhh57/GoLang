package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type userController struct {
	userService UserService
}

func NewUserController(userService UserService) *userController {
	return &userController{userService}
}

func (h *userController) CreateUser(c *gin.Context) {
	log.Default().Printf("========== Start-Method-CreateUser ==========")
	var userRequest UserRequest
	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		log.Default().Printf("========== start err != nil ==========")
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			log.Default().Printf("========== start-check-error ==========")
			errorMessage := fmt.Sprintf("Error on filed %s conditon: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			log.Default().Printf("========== end-check-error ==========")
			return
		}
	}
	user, err := h.userService.Create(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

	log.Printf("data: %s", h)
	log.Default().Printf("========== End-Method-CreateUser ==========")
}
