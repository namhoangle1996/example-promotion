package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) GetListLuckyMoneyMessage(
	ctx context.Context,
	in *service_promotion.GetListLuckyMoneyMessageRequest,
	out *service_promotion.GetListLuckyMoneyMessageResponse,
) error {
	// implement me pls
	return nil
}
