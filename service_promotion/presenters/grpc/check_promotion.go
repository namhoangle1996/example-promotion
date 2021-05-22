package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) CheckPromotion(
	ctx context.Context,
	in *service_promotion.CheckPromotionRequest,
	out *service_promotion.CheckPromotionResponse,
) error {

	voucher_detail, _, err := presenter.user_interactor.CheckUserCanUseVoucher(in.UserId, in.Code)

	if err != nil {
		return err
	}

	out.Voucher = ConvertEntitiesVoucher2DTO(voucher_detail)

	return nil
}
