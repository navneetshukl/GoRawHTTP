package main

import (
	"log"

	"github.com/navneetshukl/gorawhttp/internal/rawHttp"
)

func CheckHealth(ctx *rawHttp.Context) {
	log.Println("CheckHealth : Method is ", ctx.Method)
	log.Println("Check Health : Path is ", ctx.Path)
	ctx.String(200, "workingfine")
}

func JsonResponse(ctx *rawHttp.Context) {
	ctx.JSON(400, rawHttp.H{
		"firstName": "navneet",
		"lastName":  "shukla",
	})
}

func main() {
	router := rawHttp.NewRouter()
	router.GET("/", CheckHealth)
	router.GET("/json", JsonResponse)
	router.Run()
}
