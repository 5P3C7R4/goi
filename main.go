package main

import (
	"fmt"
	"goi/pkg/goi"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

// type Login struct {
// 	UserName string `goi:"userName" json:"userName"`
// 	Password string `goi:"password" json:"password"`
// }

func handleTest(ctx *gin.Context) {
	var data any

	ctx.BindJSON(&data)
	// err := goi.Schema(map[string]any{
	// 	"userName": goi.String().Required().Min(1).LowerCase().Trim(),
	// 	"password": goi.String().Required().Min(8),
	// }).Validate(&data)

	err := goi.String().Required().LowerCase().Validate(&data)

	// err := goi.Number().Required().Validate(&data)
	if err != nil {
		ctx.JSON(400, map[string]string{"code": err.Error()})
		return
	}
	// var newData Login
	// castedData := data.(map[string]any)
	// goi.DecodeCustom(castedData, &newData)
	ctx.JSON(200, data)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	r := gin.Default()
	r.POST("test", handleTest)

	fmt.Println("Listening and serving on 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor", err)
	}
}
