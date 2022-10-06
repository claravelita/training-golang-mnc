package dtos

import "time"

type (
	OrderRequest struct {
		CustomerName string    `json:"customer_name" validate:"required"`
		OrderedAt    time.Time `json:"ordered_at" validate:"required"`

		Items []ItemRequest `json:"items"`
	}
)
