package entity

type Service struct {
	Id          int64  `json:"id"`
	Icon        string `json:"icon"`
	BgColor     string `json:"bg_color"`
	ServiceName string `json:"service_name"`
}
