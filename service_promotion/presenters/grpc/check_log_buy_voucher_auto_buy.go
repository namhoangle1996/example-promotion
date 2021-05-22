package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) CheckLogBuyVoucherAutoBuy(
	ctx context.Context,
	in *service_promotion.CheckLogBuyVoucherAutoBuyRequest,
	out *service_promotion.CheckLogBuyVoucherAutoBuyResponse,
) error {
	// implement me pls
	return nil
}
