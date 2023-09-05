package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go-api-mysql/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "hainh:Admin123@tcp(localhost:3306)/mysql-go?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Default().Fatal("Database connection error")
	}
	log.Default().Printf("Database connection successd")
	log.Default().Printf("date", timestamp.Timestamp{})
	db.AutoMigrate(&user.User{})

	userRepository := user.NewRepository(db)

	userService := user.NewUserService(userRepository)

	userController := user.NewUserController(userService)

	//userRequest := user.UserRequest{
	//	Name:       "Hai",
	//	DayOfBirth: "05/07/2023",
	//	Phone:      "0346135365",
	//	Email:      "nguyenhoanghai@vietbank.com.vn",
	//	UserName:   "hainh",
	//	PassWord:   "hainh123",
	//}
	//
	//userService.Create(userRequest)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.POST("/login", userController.CreateUser)

	router.Run()

}
