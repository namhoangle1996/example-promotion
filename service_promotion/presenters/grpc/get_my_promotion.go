package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) GetMyPromotion(
	ctx context.Context,
	in *service_promotion.GetMyPromotionRequest,
	out *service_promotion.GetMyPromotionResponse,
) error {
	var dto_vouchers = []*service_promotion.VoucherUserDetail{}
	aggs_voucher, err := presenter.user_interactor.GetMyVoucher(0, in.UserId)

	if err != nil {
		return err
	}
	for _, v := range aggs_voucher {

		dto_vouchers = append(dto_vouchers, &service_promotion.VoucherUserDetail{
			VoucherDetail: &service_promotion.VoucherDetail{
				Voucher:  ConvertEntitiesVoucher2DTO(&v.VoucherDetail.Voucher),
				Merchant: ConvertEntitiesMerchant2DTO(&v.VoucherDetail.MerchantAccount),
			},
			VoucherUser: &service_promotion.PromotionUser{
				Id:          v.PromotionUser.Id,
				IdVoucher:   v.PromotionUser.Id,
				IdUser:      v.PromotionUser.Id,
				Amount:      int64(v.PromotionUser.Amount),
				Total:       v.PromotionUser.Total,
				TotalFreeze: v.PromotionUser.TotalFreeze,
				FreezeIds:   v.PromotionUser.FreezeIds,
				Code:        v.PromotionUser.Code,
				Status:      v.PromotionUser.Status,
				CreatedAt:   v.PromotionUser.CreatedAt.Unix(),
				LastUsed:    v.PromotionUser.LastUsed.Unix(),
				LastBuy:     v.PromotionUser.LastBuy.Unix(),
				UpdatedAt:   v.PromotionUser.UpdatedAt.Unix(),
				V:           v.PromotionUser.V,
			},
		})
	}
	out.Vouchers = dto_vouchers

	return nil
}
