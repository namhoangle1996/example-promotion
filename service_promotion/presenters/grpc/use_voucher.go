package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) UseVoucher(
	ctx context.Context,
	in *service_promotion.UseVoucherRequest,
	out *service_promotion.UseVoucherResponse,
) error {

	v, err := presenter.user_interactor.UseVoucher(in, out)
	if err != nil {
		return err
	}

	out.Voucher = &service_promotion.VoucherDetail{
		Voucher:              ConvertEntitiesVoucher2DTO(v),
		Merchant:             nil,
	}

	return err
}
