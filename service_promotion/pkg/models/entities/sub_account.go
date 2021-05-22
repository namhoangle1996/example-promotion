package entities

type ESubAccount struct {
	SubAccountId            string `bson:"sub_account_id"`
	PromotionName           string `bson:"promotion_name,omitempty"`
	PromotionCode           string `bson:"promotion_code,omitempty"`
	Amount                  int64  `bson:"amount,omitempty"`
	GpayID                  string `json:"gpay_id,omitempty"`
	MerchantID              string `bson:"merchant_id,omitempty"`
	TotalPromotionInventory int64  `bson:"total_promotion_inventory,omitempty"`
	MaxPaidAmount           int64  `bson:"max_paid_amount,omitempty"`
}

