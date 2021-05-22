package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) GetListLuckyMoneyAmount(
	ctx context.Context,
	in *service_promotion.GetListLuckyMoneyAmountRequest,
	out *service_promotion.GetListLuckyMoneyAmountResponse,
) error {
	// implement me pls
	return nil
}
