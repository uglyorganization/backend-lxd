package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

func main() {
	Start()
}

// Start echo server
func Start() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger = logger.With("env", "dev")
	config := slogecho.Config{
		WithRequestID:    true,
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
	}

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(slogecho.NewWithConfig(logger, config))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	g := e.Group("/v1")
	g.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == "valid-key", nil
	}))
	g.GET("/protected", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
