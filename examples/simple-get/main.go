package main

import (
	"fmt"

	"github.com/navneetshukl/gorawhttp/internal/middlewares"
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

func main() {
	router := rawHttp.NewRouter()
	router.UseMiddleware(middlewares.Logger)
	//router.GET("/", CheckHealth)
	router.GET("/json", JsonResponse)
	router.Run()
}
