package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) CheckUserLuckyMoney(
	ctx context.Context,
	in *service_promotion.CheckUserLuckyMoneyRequest,
	out *service_promotion.CheckUserLuckyMoneyResponse,
) error {
	// implement me pls
	return nil
}
