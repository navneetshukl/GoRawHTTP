package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/gorawhttp/internal/rawHttp"
)

func App(ctx *gin.Context) {
	time.Sleep(10 * time.Second)
	fmt.Println("1")
	ctx.JSON(200, gin.H{
		"name": "navneet",
	})
	fmt.Println("2")
}

func main() {
	//server.ListenAndServe()

	// router := gin.New()
	// router.Use(gin.Logger())
	// router.GET("/app", App)
	// router.Run()

	router:=rawHttp.NewRouter()
	router.Run()
}
