package entity

type User struct {
	Id          int    `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Picture     string `json:"picture"`
	Address     string `json:"address"`
	PinPoint    string `json:"pin_point"`
}
