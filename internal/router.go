package internal

import (
	"fiber/internal/controller"
	"fiber/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"time"
)

func NewRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       false, // CPU亲和性，Goland断点调试必须设为false !!!!!!
		CaseSensitive: true,
		StrictRouting: true, // 例如"/foo"和"/foo/" 认为是不同的路由
		ServerHeader:  "Fiber-Layout",
		AppName:       "v0.0.1",
	})

	addMiddlewareDemo(app) // 添加中间件

	app.Get("/", controller.SayHello)
	app.Get("/error", controller.GetError)
	app.Get("/error/wrap", controller.WrapErrorLog)
	app.Post("/body", controller.GetPostBody)
	app.Get("/panic", controller.GetPanic)
	app.Get("/args/:arg", controller.GetArg)
	app.Post("/bind/:arg", controller.BindJson)
	app.Get("/fast", controller.GetFastHttpRequest)
	app.Post("/fast", controller.PostFastHttpRequest)
	app.Use(func(c *fiber.Ctx) error { // 404，必须写在所有路由后面
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "no such route"})
	})

	return app
}

func addMiddlewareDemo(app *fiber.App) {
	app.Use(middleware.Recovery())          // panic恢复
	app.Use(middleware.GenerateRequestID()) // 请求唯一id
	app.Use(middleware.RequestLogger())     // 请求日志
	// app.Use(middleware.GetSwaggerUi("/v1/swagger")) // swagger文档
	app.Use(pprof.New(pprof.Config{Prefix: "/v1"})) // pprof接口暴露，接口："/v1/debug/pprof/"

	// metrics信息，ui页面"/metrics"，api访问时请求头携带Accept: application/json即可
	app.Get("/metrics", monitor.New(monitor.Config{
		Title:   "FiberLayout Metrics Page",
		Refresh: time.Second * 5, // refresh interval
		APIOnly: false,           // 允许通过ui界面访问
	}))
}
