package middleware

import (
	"boilerplate/internal/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logger.Log
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		if lastErr := ctx.Errors.Last(); lastErr != nil {
			log.WithFields(map[string]any{
				"METHOD":    reqMethod,
				"URI":       reqUri,
				"STATUS":    statusCode,
				"LATENCY":   latencyTime,
				"CLIENT_IP": clientIP,
			}).Error(lastErr.Err.Error())
			return
		}

		log.WithFields(map[string]any{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}).Infof("REQUEST %s %s SUCCESS", reqMethod, reqUri)
	}
}
