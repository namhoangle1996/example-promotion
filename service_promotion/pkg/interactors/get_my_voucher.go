package interactors

import (
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/aggregates"
)

func (i InteractorImpl) GetMyVoucher(offset int64, user_id string) ([]*aggregates.VoucherUserDetail, error) {
	var aggs_voucher = []*aggregates.VoucherUserDetail{}
	entities_popup_user, err := i.IUserHasVoucher.Get(user_id, offset)
	if err != nil {
		return nil, err
	}

	for _, entity_popup_user := range entities_popup_user {
		agg_voucher, err := i.GetDetail(entity_popup_user.IdVoucher)

		if err == nil {
			aggs_voucher = append(aggs_voucher, &aggregates.VoucherUserDetail{
				VoucherDetail: *agg_voucher,
				PromotionUser: *entity_popup_user,
			})
		}
	}

	return aggs_voucher, nil
}
