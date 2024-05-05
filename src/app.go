package src

import (
	"fmt"

	"github.com/bhagas/go-svc-echo/config"
	"github.com/bhagas/go-svc-echo/src/users"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type server struct {
	httpServer *echo.Echo
	cfg        config.Config
}

type Server interface {
	Run()
}

func InitServer(cfg config.Config) Server {
	e := echo.New()

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.HideBanner = true
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"*"},
	}))

	// Routes
	// e.GET("/users", users.GetAllUsers)
	// e.POST("/users", createUser)
	// e.GET("/users/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	// Start server
	// e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port())))
	return &server{
		httpServer: e,
		cfg:        cfg,
	}

}

func (c *server) Run() {
	//health check
	userGroup := c.httpServer.Group(`/users`)
	users.Pasang(userGroup, c.cfg)
	if err := c.httpServer.Start(fmt.Sprintf(":%d", c.cfg.Port())); err != nil {
		c.httpServer.Logger.Fatal(err)
	}
}
