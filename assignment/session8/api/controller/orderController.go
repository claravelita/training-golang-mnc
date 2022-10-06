package controller

import (
	"net/http"
	"strconv"

	"github.com/claravelita/training-golang-mnc/assignment/session8/common/command"
	"github.com/claravelita/training-golang-mnc/assignment/session8/dtos"
	"github.com/claravelita/training-golang-mnc/assignment/session8/usecase"
	"github.com/labstack/echo/v4"
)

type orderController struct {
	usecase usecase.OrderUsecase
}

func NewOrderController(usecase usecase.OrderUsecase) *orderController {
	return &orderController{usecase: usecase}
}

func (c *orderController) Route(group *echo.Group) {
	group.POST("/orders", c.Create)
	group.GET("/orders", c.GetAll)
	group.GET("/orders/:id", c.GetById)
	group.DELETE("/orders/:id", c.Delete)
	group.PUT("/orders/:id", c.Update)
}

func (c *orderController) Create(ctx echo.Context) error {
	request := dtos.OrderRequest{}
	err := command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.CreateOrderCommand(request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

func (c *orderController) GetAll(ctx echo.Context) error {
	responses, err := c.usecase.FindAllOrder()
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

func (c *orderController) GetById(ctx echo.Context) error {
	parameter := ctx.Param("id")
	id, err := strconv.Atoi(parameter)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, command.BadRequestResponses("parameter order_id should be integer"))
	}
	responses, err := c.usecase.FindOrderByID(id)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

func (c *orderController) Delete(ctx echo.Context) error {
	parameter := ctx.Param("id")
	id, err := strconv.Atoi(parameter)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, command.BadRequestResponses("parameter order_id should be integer"))
	}
	responses, err := c.usecase.DeleteOrderByID(id)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

func (c *orderController) Update(ctx echo.Context) error {
	parameter := ctx.Param("id")
	id, err := strconv.Atoi(parameter)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, command.BadRequestResponses("parameter order_id should be integer"))
	}

	request := dtos.OrderRequest{}
	err = command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}
	responses, err := c.usecase.UpdateOrderByID(id, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
