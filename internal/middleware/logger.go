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

		// Execute next handlers
		ctx.Next()

		latency := time.Since(start)
		status := ctx.Status
		method := ctx.Method
		path := ctx.Path

		// ---------- STATUS COLOR + ICON ----------
		var statusColor *color.Color
		var statusIcon string

		switch {
		case status >= 500:
			statusColor = color.New(color.FgRed, color.Bold)
			statusIcon = "🔴"
		case status >= 400:
			statusColor = color.New(color.FgYellow, color.Bold)
			statusIcon = "🟡"
		case status >= 300:
			statusColor = color.New(color.FgCyan, color.Bold)
			statusIcon = "🔵"
		default:
			statusColor = color.New(color.FgGreen, color.Bold)
			statusIcon = "🟢"
		}

		// ---------- METHOD COLOR ----------
		methodColor := color.New(color.FgMagenta, color.Bold)

		// ---------- PATH COLOR ----------
		pathColor := color.New(color.FgWhite)

		// ---------- LATENCY COLOR ----------
		latencyColor := color.New(color.FgBlue)

		// ---------- FORMAT VALUES ----------
		statusStr := statusColor.Sprintf("%s %3d", statusIcon, status)
		methodStr := methodColor.Sprintf("%-7s", method)
		pathStr := pathColor.Sprintf("%-30s", path)
		latencyStr := latencyColor.Sprintf("%8s", latency)

		// ---------- FINAL OUTPUT ----------
		fmt.Printf(
			"%s │ %s │ %s │ %s\n",
			statusStr,
			methodStr,
			pathStr,
			latencyStr,
		)
	}
}
