package model

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id" validate:"required"`
}

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" validate:"required"`
	PageSize int `json:"pageSize" form:"pageSize" validate:"required"`
}

