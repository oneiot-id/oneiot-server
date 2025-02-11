package entity

type Review struct {
	Id            int     `json:"id"`
	UserId        int     `json:"user_id"`
	TransactionId int     `json:"transaction_id"`
	Rating        float32 `json:"rating"`
	Commentary    string  `json:"commentary"`
}
