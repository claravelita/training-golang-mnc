package domain

type Item struct {
	ItemID      int    `json:"item_id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	OrderID     int    `json:"order_id" gorm:"NOT NULL"`
	ItemCode    string `json:"item_code" gorm:"DEFAULT NULL;type:varchar(45)"`
	Description string `json:"description" gorm:"DEFAULT NULL;type:varchar(45)"`
	Quantity    int    `json:"quantity"`
}
