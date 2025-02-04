package entity

type Buyer struct {
	Id              int64  `json:"id"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	FullAddress     string `json:"full_address"`
	AdditionalNotes string `json:"additional_notes"`
}
