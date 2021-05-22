package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) CheckLogsUserLuckyMoney(
	ctx context.Context,
	in *service_promotion.CheckLogsUserLuckyMoneyRequest,
	out *service_promotion.CheckLogsUserLuckyMoneyResponse,
) error {
	// implement me pls
	return nil
}
