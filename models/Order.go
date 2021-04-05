package models

import "gorm.io/gorm"

type Order struct {
	Id         uint        `json:"id"`
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"`
	Email      string      `json:"email"`
	Total      float64     `json:"total" gorm:"-"`
	UpdatedAt  string      `json:"updated_at"`
	CreatedAt  string      `json:"deleted_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float64 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)
	return total
}

func (order *Order) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order
	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)
	for i, _ := range orders {
		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName

		var total float64
		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float64(orderItem.Quantity)
		}

		orders[i].Total = total
	}
	return orders
}
