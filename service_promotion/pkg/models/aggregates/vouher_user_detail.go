package aggregates

import "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"

type VoucherUserDetail struct {
	VoucherDetail
	entities.PromotionUser
}
