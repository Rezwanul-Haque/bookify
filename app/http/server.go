package http

import (
	container "bookify/app"
	"bookify/app/http/middlewares"
	"bookify/infra/config"
	"bookify/infra/logger"
	"context"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"time"
)

func Start() {
	e := echo.New()
	lc := logger.Client()

	if err := middlewares.Attach(e, lc); err != nil {
		logger.Client().Error("error occurred when attaching middlewares", err)
		os.Exit(1)
	}

	// routes for documentation
	dg := e.Group("docs")
	dg.GET("/swagger", echo.WrapHandler(middlewares.SwaggerDocs()))
	dg.GET("/redoc", echo.WrapHandler(middlewares.ReDocDocs()))
	dg.GET("/rapidoc", echo.WrapHandler(middlewares.RapiDocs()))
	e.File("/swagger.yaml", "./swagger.yaml")

	container.Init(e.Group("api"), lc)

	port := config.App().Port

	// start http server
	go func() {
		e.Logger.Fatal(e.Start(":" + port))
	}()

	// graceful shutdown
	GracefulShutdown(e, lc)
}

// GracefulShutdown server will gracefully shut down within 5 sec
func GracefulShutdown(e *echo.Echo, lc logger.LogClient) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	lc.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = e.Shutdown(ctx)
	lc.Info("server shutdowns gracefully")
}
