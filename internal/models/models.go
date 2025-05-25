package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`     
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Role	  string `gorm:"default:'user'" json:"role"`
	Orders    []Order 
}


type Product struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`     
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Name        string  `gorm:"not null" json:"name"`
	Description string `json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int     `gorm:"not null" json:"stock"`

	// Relación con imágenes del producto (opcional)
	Images      []ProductImage  `gorm:"foreignKey:ProductID" json:"images,omitempty"`

	CategoryID  uint `json:"category_id"`
	Category    Category `json:"category"`

	Weight      float64         `gorm:"type:decimal(10,2)" json:"weight,omitempty"`
	Dimensions  string          `gorm:"size:50" json:"dimensions,omitempty"`
}

type ProductImage struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	ProductID  uint           `json:"productId"`
	ImageURL   string         `gorm:"not null" json:"imageUrl"`
	IsPrimary  bool           `gorm:"default:false" json:"isPrimary"`
	AltText    string         `gorm:"size:255" json:"altText,omitempty"`
}

type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`     
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `json:"description"`
	Products    []Product
	ParentID  *uint     `json:"parent_id" gorm:"index"`
    Parent    *Category `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
    Children  []Category `json:"children,omitempty" gorm:"foreignKey:ParentID"`
}

type Order struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`     
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	UserID      uint
	User        User
	Total       float64 `gorm:"not null"`
	Status      string  `gorm:"not null"`
	PaymentID   string
	ShippingID  string
}


type RegisterRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}