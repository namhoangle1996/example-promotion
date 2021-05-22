package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) BuyVoucher(
	ctx context.Context,
	in *service_promotion.BuyVoucherRequest,
	out *service_promotion.BuyVoucherResponse,
) error {
	return presenter.user_interactor.BuyVoucher(in)
}
