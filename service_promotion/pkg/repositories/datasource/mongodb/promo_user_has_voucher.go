package mongodb

import "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"

type IUserHasVoucher interface {
	FindByCode(user_id, code string) (*entities.PromotionUser, error)
	FindById(user_id, id string) (*entities.PromotionUser, error)
	Create(user_voucher entities.PromotionUser) error
	Update(user_voucher entities.PromotionUser) error
	DecrementTotalVoucher(id, user_id string, total int64) error
	IncrementTotalVoucher(id, user_id string, total int64) error
	Get(user_id string, offset int64) ([]*entities.PromotionUser, error)
}
