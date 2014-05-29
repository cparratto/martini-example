package models

import (
	"github.com/joiggama/hood"
)

type Products struct {
	Id        hood.Id      `json:"id"`
	Code      string       `json:"code"       binding:"required"`
	Name      string       `json:"name"       binding:"required"`
	CreatedAt hood.Created `json:"created_at"`
	UpdatedAt hood.Updated `json:"updated_at"`
}

type ProductsCollection struct {
	Elements []Products `json:"products"`
}
