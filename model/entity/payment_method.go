package entity

type PaymentMethod struct {
	Id      int64  `json:"id"`
	Name    string `json:"payment_name"`
	Number  string `json:"payment_number"`
	Logo    string `json:"payment_logo"`
	Acronym string `json:"payment_acronym"`
}
