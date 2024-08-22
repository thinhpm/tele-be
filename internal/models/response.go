package models

type CustomResponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}
