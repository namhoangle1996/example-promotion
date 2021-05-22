package mongodb

import (
	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

type IVoucher interface {
	FindById(id string) (*entities.Voucher, error)
	FindByCode(code string) (*entities.Voucher, error)
	GetList(in *service_promotion.GetListPromotionRequest) ([]*entities.Voucher, error)
	IncrementTotal(voucher_id string, total uint64) error
	DecrementTotal(voucher_id string, total uint64) error
}
