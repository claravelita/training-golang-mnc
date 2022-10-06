package dtos

type (
	ItemRequest struct {
		ItemID      int    `json:"item_id"`
		ItemCode    string `json:"item_code" validate:"required"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}
)
