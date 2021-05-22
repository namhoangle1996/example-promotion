package interactors

import (
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/aggregates"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

func (i InteractorImpl) GetDetail(voucher_id string) (*aggregates.VoucherDetail, error) {
	entity_voucher, err := i.IVoucher.FindById(voucher_id)

	if err != nil {
		return nil, err
	}

	var merchantE *entities.Merchant
	merchantE, err = i.IServiceUser.FindMerchantByCode(entity_voucher.MerchantCode)

	if err != nil {
		merchantE = &entities.Merchant{
			Logo: "https://cms-admin.g-pay.xyz/img/logo-gpay.png", // fix cứng nếu ko tìm thấy merchant
		}
	}

	return &aggregates.VoucherDetail{
		Voucher:         *entity_voucher,
		MerchantAccount: *merchantE,
	}, nil
}
