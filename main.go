package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userByEmail, err := userRepository.FinByEmail("com.johan@gmail,com")
	if err != nil {
		fmt.Println(err.Error())

	}

	if userByEmail.ID == 0 {
		fmt.Println("user tidak di temukan ")
	} else {
		fmt.Println(userByEmail.Name)
	}

	userHandler := handler.NewUserHanler(userService)

	router := gin.Default()

	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}
