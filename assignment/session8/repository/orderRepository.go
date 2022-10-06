package repository

import (
	"errors"

	"github.com/claravelita/training-golang-mnc/assignment/session8/domain"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order domain.Order) (domain.Order, error)
	GetAllOrder() (order []domain.Order, err error)
	GetOrderById(id int) (order *domain.Order, err error)
	DeleteOrder(id int) (err error)
	UpdateOrder(id int, order domain.Order) (domain.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return orderRepository{
		db,
	}
}

func (r orderRepository) CreateOrder(order domain.Order) (domain.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (r orderRepository) GetAllOrder() (order []domain.Order, err error) {
	err = r.db.Preload("Items").Find(&order).Error
	if err != nil {
		return []domain.Order{}, err
	}

	return order, nil
}

func (r orderRepository) GetOrderById(id int) (order *domain.Order, err error) {
	err = r.db.Preload("Items").First(&order, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return order, err
}

func (r orderRepository) DeleteOrder(id int) (err error) {
	return r.db.Delete(&domain.Order{}, id).Error
}

func (r orderRepository) UpdateOrder(id int, order domain.Order) (domain.Order, error) {
	err := r.db.Where("order_id = ?", id).Updates(&order).Error
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
