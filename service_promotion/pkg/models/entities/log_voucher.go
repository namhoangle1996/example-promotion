package entities

import "time"

type LogUserVoucher struct {
	Id        string    `json:"_id" bson:"_id"`
	IdVoucher string    `json:"id_voucher" bson:"id_voucher"`
	IdUser    string    `json:"id_user" bson:"id_user"`
	Action    string    `json:"action"`
	GpayId    string    `json:"gpay_id" bson:"gpay_id"`
	State     string    `json:"state" bson:"state"`
	TraceID   string    `bson:"trace_id"`
	Total     int64     `json:"total"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
