package interactors

import "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/aggregates"

func (i InteractorImpl) GetDetailByCode(code string) (*aggregates.VoucherDetail, error) {
	entity_voucher, err := i.IVoucher.FindByCode(code)

	if err != nil {
		return nil, err
	}

	merchant, err := i.IServiceUser.FindMerchantById(entity_voucher.MerchantID)

	if err != nil {
		return nil, err
	}

	return &aggregates.VoucherDetail{
		Voucher:         *entity_voucher,
		MerchantAccount: *merchant,
	}, nil
}
