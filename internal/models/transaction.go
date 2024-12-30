package models

import "github.com/go-playground/validator/v10"

type TransactionRequest struct {
	Reference string  `json:"reference" validate:"required"`
	Amount    float64 `json:"amount" validate:"required"`
}

func (I *TransactionRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

type TransactionResponse struct {
	Balance float64 `json:"balance"`
}
