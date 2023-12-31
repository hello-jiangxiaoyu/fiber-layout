package middleware

import (
	"errors"
	"fiber/internal/endpoint/resp"
	"fiber/pkg/global"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

func stackInfo(msg string, err any, skip int, fullStack bool) string {
	pwd, _ := os.Getwd()
	pwd = strings.ReplaceAll(pwd, `\`, "/") // handle windows path
	res := fmt.Sprintf("\n[Recovery] panic recovered: %s\n[Error] %v", msg, err)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		if fullStack || pwd == "" || strings.Contains(file, pwd) { // only about quick auth source files
			res += fmt.Sprintf("\n\t%s:%d %s", file, line, runtime.FuncForPC(pc).Name())
		}
	}
	return res + "\n"
}

func Recovery() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") || strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				req := fmt.Sprintf("requestID:%v method:%s path:%s", c.Locals(fiber.HeaderXRequestID), c.Method(), c.Path())
				global.Log.Error(stackInfo(req, err, 3, global.Config.Log.IsFullStack))
				if !brokenPipe {
					_ = resp.ErrorPanic(c)
				}
				zapLog(c, time.Now(), err)
			}
		}()
		return c.Next()
	}
}
