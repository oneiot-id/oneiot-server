package entity

import "time"

type Order struct {
	Id            int64
	UserId        int64
	BuyerId       int64
	OrderDetailId int64
	IsActive      bool
	CreatedAt     time.Time
}
