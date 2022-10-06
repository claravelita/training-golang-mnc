package dtos

import "time"

type (
	OrderRequest struct {
		CustomerName string    `json:"customer_name" validate:"required" example:"Tom Jerry"`
		OrderedAt    time.Time `json:"ordered_at" validate:"required" example:"2022-10-06T22:15:00+07:00"`

		Items []ItemRequest `json:"items"`
	}

	OrderUpdateRequest struct {
		CustomerName string    `json:"customer_name" validate:"required" example:"Spyke Tyke"`
		OrderedAt    time.Time `json:"ordered_at" validate:"required" example:"2019-10-09T22:15:00+07:00"`

		Items []ItemUpdateRequest `json:"items"`
	}
)
