package main

import (
	"fmt"

	"github.com/navneetshukl/gorawhttp/internal/rawHttp"
)

type Data struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SendName(ctx *rawHttp.Context) {
	data := &Data{}

	ctx.DecodeBodyStruct(data)
	fmt.Println("name is ", data.Name)
	fmt.Println("email is ", data.Email)

	ctx.JSON(200, rawHttp.H{
		"msg": "data read successfully",
	})

}

func main() {
	router := rawHttp.NewRouter()
	router.POST("/", SendName)
	router.Run()
}
