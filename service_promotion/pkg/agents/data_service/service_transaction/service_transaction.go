package service_transaction

import (
	"github.com/micro/go-micro/v2"
	"gitlab.com/wallet-gpay/service-promotion/proto/service_transaction"
)

type repositoryImpl struct {
	Service service_transaction.Service
}

// FindDetailById

func NewRepositoryImpl(serviceName string) *repositoryImpl {
	service := micro.NewService()
	service.Init()

	serviceCore := service_transaction.NewService(serviceName, service.Client())

	return &repositoryImpl{
		Service: serviceCore,
	}
}
