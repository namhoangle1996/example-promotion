package entities

import "time"

type Voucher struct {
	Id                 string        `json:"id" bson:"_id"`
	Title              string        `json:"title" bson:"title"`
	Thumbnail          string        `json:"thumbnail" bson:"thumbnail"`
	Rule               string        `json:"rule" bson:"rule"`
	Description        string        `json:"description" bson:"description"`
	ShortDescription   string        `json:"short_description" bson:"short_description"`
	Code               string        `json:"voucher_code" bson:"voucher_code"`
	Amount             int           `json:"amount" bson:"amount"`
	Type               string        `json:"type" bson:"type"`
	Total              int64         `json:"total" bson:"total"`
	TotalFreeze        int64         `json:"total_freeze" bson:"total_freeze"`
	FreezeIds          []string      `json:"freeze_ids" bson:"freeze_ids"`
	TotalCanBuy        int64         `json:"total_can_buy" bson:"total_can_buy"`
	TotalCanBuyPerDay  int64         `json:"total_can_buy_per_day" bson:"total_can_buy_per_day"`
	TotalCanUsePerDay  int64         `json:"total_can_use_per_day" bson:"total_can_use_per_day"`
	TotalCanUsePerWeek int64         `json:"total_can_use_per_week" bson:"total_use_per_week"`
	TotalCanUse        int64         `json:"total_can_use" bson:"total_can_use"`
	DiscountAmount     int64         `json:"discount_amount" bson:"discount_amount"`
	Percent            float64       `json:"discount_percent" bson:"discount_percent"`
	MerchantID         string        `json:"merchant_id" bson:"id_merchant"`
	MerchantCode       string        `json:"merchant_code" bson:"merchant_code"`
	MaxAmount          int64         `json:"max_discount_amount"  bson:"max_discount_amount"`
	Status             string        `json:"status" bson:"status"`
	Show               bool          `json:"show" bson:"show"`
	AllowAll           bool          `json:"allow_all" bson:"allow_all"`
	ShowInMyVoucher    bool          `json:"show_in_my_voucher" bson:"show_in_my_voucher"`
	StartedAt          int64         `json:"started_at" bson:"started_at"`
	ExpiredAt          int64         `json:"expired_at" bson:"expired_at"`
	CreatedAt          time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at" bson:"updated_at"`
	V                  int64         `json:"_v" bson:"_v"`
	MerchantIDs        []string      `json:"merchant_ids" bson:"merchant_ids"`
	TotalCanBuyPer     string        `json:"total_can_buy_per" bson:"total_can_buy_per"`
	TotalCanUsePer     string        `json:"total_can_use_per" bson:"total_can_use_per"`
	MinTransValue      int64         `json:"min_trans_value" bson:"min_trans_value"` //giá trị giao dịch tối thiểu để áp dụng km
	SourceUser         string        `json:"source_user" bson:"source_user"`         // nguồn user GAPO | "app"
	SourceOfFund       string        `json:"source_of_fund" bson:"source_of_fund"`
	SourceOfFunds      SourceOffunds `json:"source_of_funds" bson:"source_of_funds"`
}

type SourceOffunds []string

func (this SourceOffunds) IsContains(source string) bool {
	for _, v := range this {
		if v == source {
			return true
		}
	}
	return false
}
