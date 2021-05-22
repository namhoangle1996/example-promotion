package aggregates

import "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"

type VoucherDetail struct {
	entities.Voucher
	MerchantAccount entities.Merchant
}
