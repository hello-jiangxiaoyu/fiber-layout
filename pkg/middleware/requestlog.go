package middleware

import (
	"fiber/pkg/global"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"os"
	"time"
)

// RequestLogger middleware to logs request/response
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next() // 继续执行后面的handler

		zapLog(c, start, err)
		return nil
	}
}

func zapLog(c *fiber.Ctx, start time.Time, err any) {
	if global.AccessLog == nil {
		fmt.Println("access log is nil")
		return
	}

	var intCode uint
	if code := c.Locals("code"); code != nil {
		intCode, _ = code.(uint)
	}
	requestID := ""
	if id := c.Locals(fiber.HeaderXRequestID); id != nil {
		requestID, _ = id.(string)
	}

	global.AccessLog.Info("",
		zap.String("ts", time.Now().Format(time.RFC3339)),
		zap.Int("pid", os.Getpid()),
		zap.String("request_id", requestID),
		zap.Int("status", c.Response().StatusCode()),
		zap.Uint("code", intCode),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.String("route", c.Route().Path),
		zap.String("query", c.Request().URI().QueryArgs().String()),
		zap.String("client_ip", c.IP()),
		zap.String("server_ip", ""),
		zap.String("host", c.Hostname()),
		zap.Duration("cost", time.Since(start)),
		zap.Int("request_length", len(c.Request().Body())),
		zap.Int("body_bytes_sent", len(c.Response().Body())),
		zap.String("referer", c.Get("Referer")),
		zap.String("proto", c.Protocol()),
		zap.String("ua", c.Get(fiber.HeaderUserAgent)),
		zap.Any("errors", err),
	)
}
