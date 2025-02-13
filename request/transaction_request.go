package request

import (
	"oneiot-server/model/dto"
	"oneiot-server/model/entity"
)

type CreateTransactionRequest struct {
	User           entity.User        `json:"user"`
	Order          entity.Order       `json:"order"`
	TransactionDto dto.TransactionDto `json:"transaction_dto"`
}

type GetTransactionRequest struct {
	User        entity.User        `json:"user"`
	Transaction entity.Transaction `json:"transaction"`
}

type GetTransactionsRequest struct {
	User entity.User `json:"user"`
}
