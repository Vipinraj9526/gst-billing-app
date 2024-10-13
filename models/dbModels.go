package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"not null;unique" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	Email     string    `gorm:"not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Product struct {
	ID                 uint      `gorm:"primaryKey"`
	ProductCode        string    `gorm:"unique;not null" json:"product_code"`
	ProductName        string    `gorm:"not null" json:"product_name"`
	ProductPrice       float64   `gorm:"not null" json:"product_price"`
	Quantity           uint64    `gorm:"not null" json:"quantity"`
	ProductDescription string    `gorm:"not null" json:"product_description"`
	ProductGST         float64   `gorm:"not null" json:"product_gst"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Bill struct {
	ID             uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	BillerUserName string     `gorm:"not null;unique" json:"biller_user_name"`
	Total          float64    `gorm:"not null" json:"total"`
	TotalTax       float64    `gorm:"not null" json:"total_tax"`
	Subtotal       float64    `gorm:"not null" json:"subtotal"`
	BillDate       time.Time  `gorm:"autoCreateTime" json:"bill_date"`
	Items          []BillItem `gorm:"foreignKey:BillID" json:"items"`
}

type BillItem struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	BillID      uint    `gorm:"not null" json:"bill_id"`
	ProductCode string  `gorm:"not null" json:"product_code"`
	ProductName string  `gorm:"-" json:"product_name"`
	Quantity    uint64  `gorm:"not null" json:"quantity"`
	Price       float64 `gorm:"not null" json:"price"`
	ProductGST  float64 `gorm:"not null" json:"product_gst"`
}
