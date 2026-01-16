package middleware

import (
	"context"
	"time"

	"github.com/navneetshukl/gorawhttp/internal/rawHttp"
)

// TimeOut middleware for request timeout
func TimeOut(duration time.Duration) rawHttp.Handler {
	return func(ctx *rawHttp.Context) {
		context, cancel := context.WithTimeout(context.Background(), duration)
		ctx.CancelFunc = cancel
		defer cancel()

		ctx.Context = context
		go func() {

			select {
			case <-ctx.Context.Done():
				ctx.String(500, "Request Time Out")
				return
			}

		}()

		ctx.Next()
	}
}
