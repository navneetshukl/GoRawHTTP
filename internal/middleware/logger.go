package middleware

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/navneetshukl/gorawhttp/internal/rawHttp"
)

func Logger() rawHttp.Handler {
	return func(ctx *rawHttp.Context) {

		start := time.Now()

		// Call next handler
		ctx.Next()

		latency := time.Since(start)
		status := ctx.Status //
		method := ctx.Method
		path := ctx.Path

		// Pick color based on status
		statusColor := color.New(color.FgGreen)
		switch {
		case status >= 500:
			statusColor = color.New(color.FgRed)
		case status >= 400:
			statusColor = color.New(color.FgYellow)
		case status >= 300:
			statusColor = color.New(color.FgCyan)
		}

		statusStr := statusColor.Sprintf("%d", status)
		methodStr := color.New(color.FgMagenta).Sprintf("%s", method)
		pathStr := color.New(color.FgWhite).Sprintf("%s", path)
		latencyStr := color.New(color.FgBlue).Sprintf("%v", latency)

		fmt.Printf("[RAWHTTP] %s | %s | %s | %s\n",
			statusStr,
			methodStr,
			pathStr,
			latencyStr,
		)
	}

}
