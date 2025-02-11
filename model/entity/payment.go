package entity

type Payment struct {
	Id               int64  `json:"id"`
	PaymentProof     string `json:"payment_proof"`
	Invoice          string `json:"invoice"`
	Paid             bool   `json:"paid"`
	PaymentMethodsId int64  `json:"payment_methods_id"`
}
