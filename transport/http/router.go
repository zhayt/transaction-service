package http

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func (s *Server) SetUpRoute() {
	v1 := s.App.Group("/api/v1")

	s.App.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	s.App.GET("/live", func(e echo.Context) error {
		return e.NoContent(http.StatusOK)
	})

	user := v1.Group("/accounts")
	user.POST("", s.handler.CreateUserAccount)
	user.PATCH("", s.handler.UpdateUserAccount)
	user.DELETE("/:id", s.handler.DeleteUserAccount)

	v1.POST("/transactions", s.handler.CreateTransaction)
}
