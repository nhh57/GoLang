package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-api-mysql/user"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	log.Default().Printf("========== start-method-login ==========")
	var user user.User
	err := c.ShouldBindJSON(&user)
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
	c.JSON(http.StatusOK, gin.H{
		"user_name": user.Name,
		"pass_word": user.PassWord,
	})
	h := gin.H{"user_name": user.Name,
		"pass_word": user.PassWord}
	log.Printf("data: %s", h)
	log.Default().Printf("========== end-method-login ==========")
}
