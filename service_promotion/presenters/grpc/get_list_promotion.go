package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) GetListPromotion(
	ctx context.Context,
	in *service_promotion.GetListPromotionRequest,
	out *service_promotion.GetListPromotionResponse,
) error {

	var return_vouchers = []*service_promotion.VoucherDetail{}

	aggregates_voucher, err := presenter.user_interactor.GetListVoucher(in)

	if err != nil {
		return err
	}

	for _, v := range aggregates_voucher {

		return_vouchers = append(return_vouchers, &service_promotion.VoucherDetail{
			Voucher:  ConvertEntitiesVoucher2DTO(&v.Voucher),
			Merchant: ConvertEntitiesMerchant2DTO(&v.MerchantAccount),
		})
	}
	out.Vouchers = return_vouchers
	return nil
}
