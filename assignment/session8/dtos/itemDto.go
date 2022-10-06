package dtos

type (
	ItemRequest struct {
		ItemCode    string `json:"item_code" validate:"required" example:"123"`
		Description string `json:"description" example:"Iphone 10X"`
		Quantity    int    `json:"quantity" example:"1"`
	}

	ItemUpdateRequest struct {
		ItemID      int    `json:"item_id,omitempty" example:"1"`
		ItemCode    string `json:"item_code" validate:"required" example:"123"`
		Description string `json:"description" example:"Iphone 10X"`
		Quantity    int    `json:"quantity" example:"10"`
	}
)
