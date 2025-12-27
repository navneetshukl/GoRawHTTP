package main

import (
	"log"

	"github.com/navneetshukl/gorawhttp/internal/rawHttp"
)


func CheckHealth(ctx *rawHttp.Context){
	log.Println("Method is ",ctx.Method)
	log.Println("Path is ",ctx.Path)
	ctx.WriteResponse("workingfine")
}

func main(){
	router:=rawHttp.NewRouter()
	router.GET("/",CheckHealth)
	router.Run()
}
