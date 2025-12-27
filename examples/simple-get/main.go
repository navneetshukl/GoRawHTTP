package main

import (
	"log"

	"github.com/navneetshukl/gorawhttp/internal/rawHttp"
)


func CheckHealth(ctx *rawHttp.Context){
	log.Println("CheckHealth : Method is ",ctx.Method)
	log.Println("Check Health : Path is ",ctx.Path)
	ctx.WriteResponse(200,"workingfine")
}

func main(){
	router:=rawHttp.NewRouter()
	router.GET("/",CheckHealth)
	router.Run()
}
