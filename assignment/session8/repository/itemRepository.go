package repository

import (
	"errors"

	"github.com/claravelita/training-golang-mnc/assignment/session8/domain"
	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item []domain.Item) ([]domain.Item, error)
	UpdateItem(ids int, item domain.Item) (domain.Item, error)
	GetItemById(id int) (item *domain.Item, err error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return itemRepository{
		db,
	}
}

func (r itemRepository) CreateItem(item []domain.Item) ([]domain.Item, error) {
	err := r.db.Create(&item).Error
	if err != nil {
		return []domain.Item{}, err
	}

	return item, nil
}

func (r itemRepository) UpdateItem(ids int, item domain.Item) (domain.Item, error) {
	err := r.db.Where("item_id = ?", ids).Updates(&item).Error
	if err != nil {
		return domain.Item{}, err
	}

	return item, nil
}

func (r itemRepository) GetItemById(id int) (item *domain.Item, err error) {
	err = r.db.First(&item, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return item, err
}
