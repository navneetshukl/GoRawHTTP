package main

import (
	"fmt"

	"github.com/navneetshukl/gorawhttp/internal/middleware"
	"github.com/navneetshukl/gorawhttp/internal/rawHttp"
)

func CheckHealth(ctx *rawHttp.Context) {
	ctx.String(200, "workingfine")
}

func JsonResponse(ctx *rawHttp.Context) {
	fmt.Println("1")
	ctx.JSON(400, rawHttp.H{
		"firstName": "navneet",
		"lastName":  "shukla",
	})
	fmt.Println("2")
}

func A(ctx *rawHttp.Context) {
	fmt.Println("Handler A")
	ctx.Next()
}
func B(ctx *rawHttp.Context) {
	fmt.Println("Handler B")
	ctx.Next()
}

func main() {
	router := rawHttp.NewRouter()
	router.UseMiddleware(middleware.Logger())
	// router.GET("/", CheckHealth)
	router.GET("/", A,B)

	router.GET("/json", JsonResponse)
	router.Run()
}
