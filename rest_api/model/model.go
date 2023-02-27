package model

type Employee struct {
	Id   int64  `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	City string `form:"city" json:"city"`
}

type Response struct {
	Status  int64      `json:"status"`
	Message string     `json:"message"`
	Data    []Employee `json:"data,omitempty"`
}
