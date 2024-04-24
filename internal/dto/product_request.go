package dto

import (
	"mime/multipart"
)

type ProductRequestForm struct {
	Media        *multipart.FileHeader `form:"media" binding:"required"`
	ProductField ProductRequestData    `form:"data" binding:"required"`
}

type EditProductRequestForm struct {
	Id           uint                   `form:"id" binding:"required"`
	Media        *multipart.FileHeader  `form:"media"`
	ProductField EditProductRequestData `form:"data"`
}

type ProductRequestData struct {
	Name        string               `json:"name" binding:"required"`
	Color       string               `json:"color" binding:"required"`
	Description string               `json:"description" binding:"required"`
	Discount    float32              `json:"discount" binding:"required"`
	Price       float32              `json:"price" binding:"required"`
	Sizes       []ProductSizeRequest `json:"sizes" binding:"required"`
}

type EditProductRequestData struct {
	Name        string                   `json:"name" binding:"required"`
	Color       string                   `json:"color" binding:"required"`
	Description string                   `json:"description" binding:"required"`
	Discount    float32                  `json:"discount" binding:"required"`
	Price       float32                  `json:"price" binding:"required"`
	Sizes       []EditProductSizeRequest `json:"sizes" binding:"required"`
}

type ProductSizeRequest struct {
	Size  uint `json:"size" binding:"required"`
	Stock uint `json:"stock" binding:"required"`
}

type EditProductSizeRequest struct {
	Id uint `json:"id" binding:"required"`
	ProductSizeRequest
}
type ProductQuery struct {
	// validate this
	Name     string  `form:"name"`
	Color    string  `form:"color"`
	Size     uint    `form:"size"`
	MinPrice float32 `form:"minPrice"`
	MaxPrice float32 `form:"maxPrice"`
	SortBy   string  `form:"sortBy"`
	Sort     string  `form:"sort"`
	Count    int     `form:"count"`
	Page     int     `form:"page"`
}
