package entities

import "time"

type PromotionUser struct {
	Id          string    `json:"_id" bson:"_id"`
	IdVoucher   string    `json:"id_voucher" bson:"id_voucher"`
	IdUser      string    `json:"id_user" bson:"id_user"`
	Amount      int       `json:"amount"`
	Total       int64     `json:"total"`
	TotalFreeze int64     `json:"total_freeze" bson:"total_freeze"`
	FreezeIds   []string  `json:"freeze_ids"  bson:"freeze_ids"`
	Code        string    `json:"code"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	LastUsed    time.Time `json:"last_used" bson:"last_used"`
	LastBuy     time.Time `json:"last_buy" bson:"last_buy"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	V           int64     `json:"_v" bson:"_v"`
}

type PromotionUserApply struct {
	ID        string `json:"_id" bson:"_id"`
	PromoID   string `json:"promo_id" bson:"promo_id"`
	UserID    string `json:"user_id" bson:"user_id"`
	Status    int    `json:"status" bson:"status"`
	StartDate int64  `json:"start_date" bson:"start_date"`
	EndDate   int64  `json:"end_date" bson:"end_date"`
	V         int64  `json:"_v" bson:"_v"`
}
