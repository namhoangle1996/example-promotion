package grpc

import (
	"context"
	"fmt"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) GetDetailPromotion(
	ctx context.Context,
	in *service_promotion.GetDetailPromotionRequest,
	out *service_promotion.GetDetailPromotionResponse,
) error {
	if in.Id != "" {
		voucher_detail, err := presenter.user_interactor.GetDetail(in.Id)
		if err != nil {
			return err
		}

		out.Voucher = &service_promotion.VoucherDetail{
			Voucher:  ConvertEntitiesVoucher2DTO(&voucher_detail.Voucher),
			Merchant: ConvertEntitiesMerchant2DTO(&voucher_detail.MerchantAccount),
		}
		return nil
	} else if in.Code != "" {
		voucher_detail, err := presenter.user_interactor.GetDetailByCode(in.Code)
		if err != nil {
			return err
		}

		out.Voucher = &service_promotion.VoucherDetail{
			Voucher:  ConvertEntitiesVoucher2DTO(&voucher_detail.Voucher),
			Merchant: ConvertEntitiesMerchant2DTO(&voucher_detail.MerchantAccount),
		}

		return nil
	}

	return fmt.Errorf("[GetDetailPromotion] cannot find voucher")
}
