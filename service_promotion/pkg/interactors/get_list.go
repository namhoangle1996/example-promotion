package interactors

import (
	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/aggregates"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

func (i InteractorImpl) GetListVoucher(in *service_promotion.GetListPromotionRequest) ([]*aggregates.VoucherDetail, error) {
	var aggregates_voucher []*aggregates.VoucherDetail

	entites_voucher, err := i.IVoucher.GetList(in)

	if err != nil {
		return nil, err
	}

	for _, entity_voucher := range entites_voucher {
		var merchantE *entities.Merchant
		merchantE, err = i.IServiceUser.FindMerchantByCode(entity_voucher.MerchantCode)
		if err != nil {
			merchantE = &entities.Merchant{
				Logo: "https://cms-admin.g-pay.xyz/img/logo-gpay.png", // fix cứng nếu ko tìm thấy merchant
			}
		}
		check_apply_for_user, _ := i.IUserApplyVoucher.CheckApplied(in.UserId, entity_voucher.Id)

		if entity_voucher.AllowAll == false && check_apply_for_user > 0 {
			if check_apply_for_user > 1 {
				entity_voucher.ExpiredAt = check_apply_for_user
			}
			aggregates_voucher = append(aggregates_voucher, &aggregates.VoucherDetail{
				Voucher:         *entity_voucher,
				MerchantAccount: *merchantE,
			})
		} else if entity_voucher.AllowAll == true {
			aggregates_voucher = append(aggregates_voucher, &aggregates.VoucherDetail{
				Voucher:         *entity_voucher,
				MerchantAccount: *merchantE,
			})
		}

	}

	return aggregates_voucher, nil
}
