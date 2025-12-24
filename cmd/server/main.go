package main

import "github.com/gin-gonic/gin"

func App(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"name":"navneet",
	})
}

func main() {
	//server.ListenAndServe()

	router:=gin.New()
	router.GET("/app",App)
	router.Run()
}
