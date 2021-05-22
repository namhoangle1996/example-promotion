package data_service

import (
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/aggregates"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

type IServiceUser interface {
	FindDetailById(id string) (*aggregates.UserDetail, error)
	FindById(id string) (*entities.User, error)
	FindMerchantByCode(code string) (*entities.Merchant, error)
	FindMerchantById(id string) (*entities.Merchant, error)
}
