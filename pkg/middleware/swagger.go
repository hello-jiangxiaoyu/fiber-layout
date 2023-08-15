package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/handlers"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const SwaggerPath = "./docs/swagger.json"

// GetSwaggerUi creates a new middleware handler
func GetSwaggerUi(basePath string) fiber.Handler {
	if _, err := os.Stat(SwaggerPath); os.IsNotExist(err) {
		panic(errors.New(fmt.Sprintf("%s file is not exist", SwaggerPath)))
	}
	specDoc, err := loads.Spec(SwaggerPath)
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(specDoc.Spec(), "", "  ")
	if err != nil {
		panic(err)
	}

	swaggerUIOpts := middleware.SwaggerUIOpts{
		BasePath: basePath,
		SpecURL:  path.Join(basePath, "swagger.json"),
		Path:     "index.html",
	}

	swaggerUiHandler := middleware.SwaggerUI(swaggerUIOpts, nil)
	specFileHandler := handlers.CORS()(middleware.Spec(basePath, b, swaggerUiHandler))
	return func(c *fiber.Ctx) error {
		handler := fasthttpadaptor.NewFastHTTPHandler(specFileHandler)
		handler(c.Context())
		return nil
	}
}
