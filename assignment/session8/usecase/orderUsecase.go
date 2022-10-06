package usecase

import (
	"github.com/claravelita/training-golang-mnc/assignment/session8/common/command"
	"github.com/claravelita/training-golang-mnc/assignment/session8/domain"
	"github.com/claravelita/training-golang-mnc/assignment/session8/dtos"
	"github.com/claravelita/training-golang-mnc/assignment/session8/repository"
)

type OrderUsecase interface {
	CreateOrderCommand(request dtos.OrderRequest) (dtos.JSONResponses, error)
	FindAllOrder() (dtos.JSONResponses, error)
	FindOrderByID(id int) (dtos.JSONResponses, error)
	DeleteOrderByID(id int) (dtos.JSONResponses, error)
	UpdateOrderByID(id int, request dtos.OrderRequest) (dtos.JSONResponses, error)
}

type orderImplementation struct {
	repo     repository.OrderRepository
	repoItem repository.ItemRepository
}

func NewOrderImplementation(repo repository.OrderRepository, repoItem repository.ItemRepository) OrderUsecase {
	return &orderImplementation{
		repo:     repo,
		repoItem: repoItem,
	}
}

func (o orderImplementation) CreateOrderCommand(request dtos.OrderRequest) (dtos.JSONResponses, error) {
	requestOrder := domain.Order{
		OrderedAt:    request.OrderedAt,
		CustomerName: request.CustomerName,
	}

	createOrder, err := o.repo.CreateOrder(requestOrder)
	if err != nil {
		return command.InternalServerResponses("Error to create order", err), err
	}

	if len(request.Items) != 0 {
		var domainItems []domain.Item
		for _, i := range request.Items {
			domainItem := domain.Item{
				OrderID:     createOrder.OrderID,
				ItemCode:    i.ItemCode,
				Description: i.Description,
				Quantity:    i.Quantity,
			}
			domainItems = append(domainItems, domainItem)
		}

		createItems, err := o.repoItem.CreateItem(domainItems)
		if err != nil {
			return command.InternalServerResponses("Error to create items", err), err
		}

		createOrder.Items = createItems
	}

	return command.SuccessResponses(createOrder), nil
}

func (o orderImplementation) FindAllOrder() (dtos.JSONResponses, error) {
	getAllOrder, err := o.repo.GetAllOrder()
	if err != nil {
		return command.InternalServerResponses("Error to get order", err), err
	}

	return command.SuccessResponses(getAllOrder), nil
}

func (o orderImplementation) FindOrderByID(id int) (dtos.JSONResponses, error) {
	getOrder, err := o.repo.GetOrderById(id)
	if getOrder == nil && err == nil {
		return command.BadRequestResponses("order_id Not found"), err
	}
	if err != nil {
		return command.InternalServerResponses("Error to get order", err), err
	}

	return command.SuccessResponses(getOrder), nil
}

func (o orderImplementation) DeleteOrderByID(id int) (dtos.JSONResponses, error) {
	getOrder, err := o.repo.GetOrderById(id)
	if getOrder == nil && err == nil {
		return command.BadRequestResponses("order_id Not found"), err
	}
	if err != nil {
		return command.InternalServerResponses("Error to get order", err), err
	}

	err = o.repo.DeleteOrder(id)
	if err != nil {
		return command.InternalServerResponses("Error to delete order", err), err
	}

	return command.SuccessResponses(nil), nil
}

func (o orderImplementation) UpdateOrderByID(id int, request dtos.OrderRequest) (dtos.JSONResponses, error) {
	getOrder, err := o.repo.GetOrderById(id)
	if getOrder == nil && err == nil {
		return command.BadRequestResponses("order_id Not found"), err
	}
	if err != nil {
		return command.InternalServerResponses("Error to get order", err), err
	}

	requestOrder := domain.Order{
		OrderedAt:    request.OrderedAt,
		CustomerName: request.CustomerName,
	}

	if len(request.Items) != 0 {
		for _, i := range request.Items {
			getItem, err := o.repoItem.GetItemById(i.ItemID)
			if getItem == nil && err == nil {
				return command.BadRequestResponses("item_id Not found"), err
			}
			if err != nil {
				return command.InternalServerResponses("Error to get item", err), err
			}

			requestUpdate := domain.Item{
				ItemCode:    i.ItemCode,
				Description: i.Description,
				Quantity:    i.Quantity,
			}

			_, err = o.repoItem.UpdateItem(i.ItemID, requestUpdate)
			if err != nil {
				return command.InternalServerResponses("Error to update item", err), err
			}
		}
	}

	updateOrder, err := o.repo.UpdateOrder(id, requestOrder)
	if err != nil {
		return command.InternalServerResponses("Error to update order", err), err
	}

	return command.SuccessResponses(updateOrder), nil
}
