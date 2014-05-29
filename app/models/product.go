package models

import (
  "github.com/eaigner/hood"
)

type Products struct {
	Id        hood.Id      `json:"id"`
	Code      string       `json:"code"       binding:"required"`
	Name      string       `json:"name"       binding:"required"`
	CreatedAt hood.Created `json:"created_at"`
	UpdatedAt hood.Updated `json:"updated_at"`
}

//type Products struct {
	//Collection []Product `json:"products"`
//}

// Timestamps

//func (p *Product) PreInsert(s gorp.SqlExecutor) error {
	//p.CreatedAt = time.Now()
	//p.UpdatedAt = p.CreatedAt
	//return nil
//}

//func (p *Product) PreUpdate(s gorp.SqlExecutor) error {
	//p.UpdatedAt = time.Now()
	//return nil
//}
