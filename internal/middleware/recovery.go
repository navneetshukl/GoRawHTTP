package middleware

import (
	"github.com/navneetshukl/gorawhttp/internal/rawHttp"
)

func Recovery() rawHttp.Handler {
	return func(ctx *rawHttp.Context) {

		defer func() {
			if r := recover(); r != nil {
				ctx.String(500, "Internal Server Error")
			}
		}()

		ctx.Next()
	}
}
