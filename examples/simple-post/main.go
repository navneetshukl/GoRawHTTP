package main

import (
	"fmt"

	"github.com/navneetshukl/gorawhttp/internal/rawHttp"
)

type Data struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SendNameStruct(ctx *rawHttp.Context) {
	data := &Data{}

	ctx.DecodeBody(data)
	fmt.Println("name is ", data.Name)
	fmt.Println("email is ", data.Email)

	ctx.JSON(200, rawHttp.H{
		"msg": "data read successfully",
	})

}

func SendNameInterface(ctx *rawHttp.Context) {
	var data interface{}

	ctx.DecodeBody(&data)
	fmt.Println("Data is ",data)

	ctx.JSON(200, rawHttp.H{
		"msg": "data read successfully",
	})

}

func main() {
	router := rawHttp.NewRouter()
	router.POST("/", SendNameStruct)
	router.POST("/interface",SendNameInterface)
	router.Run()
}
