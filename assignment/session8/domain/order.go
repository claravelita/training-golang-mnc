package domain

import "time"

type Order struct {
	OrderID      int       `json:"order_id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	CustomerName string    `json:"customer_name" gorm:"DEFAULT NULL;type:varchar(45)"`
	OrderedAt    time.Time `json:"ordered_at"`

	Items []Item `json:"items" gorm:"foreignkey:OrderID;references:OrderID"`
}
