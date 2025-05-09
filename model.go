package main

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Cnpj      string         `json:"cnpj"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Order struct {
	ID         uuid.UUID      `json:"id"`
	Customer   string         `json:"customer"`
	User_ID    uuid.UUID      `json:"user_id"` //foreign key
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	OrderItems []OrderItems   `json:"order_items"` //has many
}

type OrderItems struct {
	ID        uuid.UUID      `json:"id" gorm:"default:gen_random_uuid()"` //gorm default
	Name      string         `json:"name"`
	Price     int            `json:"price"`
	OrderID   uuid.UUID      `json:"order_id"` //foreign key
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type User struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Type      int            `json:"type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	//LastLogin
	//Group
}

type PendingService struct {
	ID          uuid.UUID      `json:"id"`
	User_ID     uuid.UUID      `json:"user_id"` //foreign key
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
