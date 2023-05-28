package server

import (
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		// First, we add the headers with need to enable CORS.
		// Make sure to adjust these headers to your needs.
		ginCtx.Header("Access-Control-Allow-Origin", "*")
		// ginCtx.Header("Access-Control-Allow-Methods", "*")
		// ginCtx.Header("Access-Control-Allow-Headers", "*")
		ginCtx.Header("Content-Type", "application/json")
		ginCtx.Header("Vary", "Origin")

		ginCtx.Next()
	}
}
