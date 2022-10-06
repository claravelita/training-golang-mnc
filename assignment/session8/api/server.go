package api

import (
	"net/http"
	"os"
	"time"

	"github.com/claravelita/training-golang-mnc/assignment/session8/api/controller"
	"github.com/claravelita/training-golang-mnc/assignment/session8/api/handler"
	customMiddleware "github.com/claravelita/training-golang-mnc/assignment/session8/api/middleware"
	"github.com/claravelita/training-golang-mnc/assignment/session8/infrastructure"
	"github.com/claravelita/training-golang-mnc/assignment/session8/repository"
	"github.com/claravelita/training-golang-mnc/assignment/session8/usecase"
	"github.com/labstack/echo/v4"
)

// Server Struct
type Server struct {
	Route *echo.Echo
}

// NewServer : Create Server Instance
func NewServer(e *echo.Echo) *Server {
	return &Server{
		e,
	}
}

func (server *Server) InitializeServer() {
	server.Route.Use(customMiddleware.UseCorsMiddleware())
	customMiddleware.UseCustomLogger(server.Route)
	handler.UseCustomValidatorHandler(server.Route)

	initDB := infrastructure.NewGormDB()
	apiGroup := server.Route.Group("/api")

	orderRepo := repository.NewOrderRepository(initDB)
	itemRepo := repository.NewItemRepository(initDB)
	orderUsecase := usecase.NewOrderImplementation(orderRepo, itemRepo)
	orderController := controller.NewOrderController(orderUsecase)
	orderController.Route(apiGroup)

	serverConfiguration := &http.Server{
		Addr:         ":" + os.Getenv("SERVER_PORT"),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))
}
