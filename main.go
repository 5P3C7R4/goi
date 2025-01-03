package main

import (
	"fmt"
	"goi/pkg/goi"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

type Login struct {
	UserName      *string  `goi:"userName" json:"userName"`
	Password      *string  `goi:"password" json:"password"`
	Terminal      *float64 `goi:"terminal" json:"terminal"`
	Latitude      *string  `goi:"latitutde" json:"latitutde"`
	Longitude     *string  `goi:"longitude" json:"longitude"`
	FlagPlatform  *float64 `goi:"flagPlatform" json:"flagPlatform"`
	XForwardedFor *string  `goi:"x-forwarded-for" json:"x-forwarded-for"`
	TokenKambi    *string  `goi:"tokenKambi" json:"tokenKambi"`
	TokenFront    *string  `goi:"tokenFront" json:"tokenFront"`
	Playerid      *string  `goi:"playerid" json:"playerid"`
}

func handleTest(ctx *gin.Context) {
	var data any

	ctx.BindJSON(&data)
	err := goi.Schema(map[string]any{
		"userName":  goi.String().Valid([]any{"5p3c7r4", "5p3c7r42"}).Required().LowerCase().Trim(),
		"password":  goi.String().Required(),
		"terminal":  goi.Number(),
		"latitutde": goi.String().Regex(`^[-+]?(?:[1-8]?\d(?:\.\d{1,10})?|90(?:\.0{1,10})?)$`),
		"longitude": goi.String().Regex(`^[-+]?(?:(?:1[0-7]\d|[1-9]?\d)(?:\.\d{1,10})?|180(?:\.0{1,10})?)$`),
		// "user-agent":      goi.String().Required(),
		"x-forwarded-for": goi.String().Custom(func(value *any, helpers *goi.Helper) any {
			fmt.Println(*value)
			// return helpers.Error("My error")
			// return "HOLAPRUEBA2"
		}),
		"flagPlatform": goi.Number().Required().Valid([]any{0, 1}),
		"tokenKambi":   goi.String().Default(""),
		"tokenFront":   goi.String().Default(""),
		"playerid":     goi.String().Default(""),
	}).Validate(&data)

	// data2 := (any)(nil)
	// err := goi.String().Default("HOLA").Validate(&data2)

	// err := goi.Number().Required().Validate(&data)
	if err != nil {
		ctx.JSON(400, map[string]string{"code": err.Error()})
		return
	}
	var newData Login
	castedData := data.(map[string]any)
	err = goi.CustomDecode(castedData, &newData)
	if err != nil {
		ctx.JSON(400, map[string]string{"code": err.Error()})
		return
	}
	ctx.JSON(200, newData)

	// ctx.JSON(200, data)
	// ctx.JSON(200, data2)
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
