package middlewares

import (
	"bookify/infra/logger"
	"github.com/labstack/echo/v4"
	"net/http"

	openMiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/labstack/echo/v4/middleware"
)

const EchoLogFormat = "time: ${time_rfc3339_nano} || ${method}: ${uri} || status: ${status} || latency: ${latency_human} \n"

// Attach middlewares required for the application, eg: sentry, newrelic etc.
func Attach(e *echo.Echo, lc logger.LogClient) error {
	// remove trailing slashes from each requests
	e.Pre(middleware.RemoveTrailingSlash())

	// echo middlewares, todo: add color to the log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: EchoLogFormat}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(context echo.Context) bool {
			return context.Request().RequestURI == "/metrics"
		},
		Level: 5,
	}))

	return nil
}

func SwaggerDocs() http.Handler {
	opts := openMiddleware.SwaggerUIOpts{
		Path:    "docs/swagger",
		SpecURL: "/swagger.yaml",
	}
	return openMiddleware.SwaggerUI(opts, nil)
}

func ReDocDocs() http.Handler {
	opts := openMiddleware.RedocOpts{
		Path:    "docs/redoc",
		SpecURL: "/swagger.yaml",
	}
	return openMiddleware.Redoc(opts, nil)
}

func RapiDocs() http.Handler {
	opts := openMiddleware.RapiDocOpts{
		Path:    "docs/rapidoc",
		SpecURL: "/swagger.yaml",
	}
	return openMiddleware.RapiDoc(opts, nil)
}
