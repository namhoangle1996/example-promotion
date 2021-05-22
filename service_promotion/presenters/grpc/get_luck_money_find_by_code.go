package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) GetLuckMoneyFindByCode(
	ctx context.Context,
	in *service_promotion.GetLuckMoneyFindByCodeRequest,
	out *service_promotion.GetLuckMoneyFindByCodeResponse,
) error {
	// implement me pls
	return nil
}
