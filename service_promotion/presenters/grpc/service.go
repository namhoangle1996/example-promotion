package grpc

import (
	"context"
	"github.com/micro/go-micro/v2"
	"gitlab.com/wallet-gpay/service-promotion/logger"
	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"

	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/interactors"
)

type Service struct {
	user_interactor *interactors.InteractorImpl
}

func (presenter *Service) ReverseWallet(c context.Context, i *service_promotion.ReverseWalletRequest, o *service_promotion.ReverseWalletRequest) error {
	*o = *i
	return presenter.user_interactor.ReverseWallet(i.TraceId)
}

// Constructor
func NewHandler(user_interactor *interactors.InteractorImpl) *Service {
	return &Service{user_interactor}
}

func (presenter *Service) StartServer() {
	service := micro.NewService(
		micro.Name("go.micro.srv.promotion"),
		micro.Version("latest"),
		micro.WrapHandler(logger.LogWrapper),
	)

	// Initialise service
	service.Init()
	err := service_promotion.RegisterPromotionServicesHandler(service.Server(), presenter)

	if err = service.Run(); err != nil {
		panic(err.Error())
	}
}
