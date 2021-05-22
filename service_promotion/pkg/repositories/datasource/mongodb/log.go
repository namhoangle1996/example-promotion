package mongodb

import "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"

type ILogVoucher interface {
	Create(voucher entities.LogUserVoucher) error
	Update(voucher entities.LogUserVoucher) error
	CountLogPerDay(user_id, voucher_id, action string) int64
	CountLogAll(user_id, voucher_id, action string) int64
	FindLogBy(trace_id string, check_reverse bool) (voucher []*entities.LogUserVoucher, err error)
	LogPerDay(user_id, voucher_id, action string, check_total int64) (err error)
	LogPerWeek(user_id, voucher_id, action string, check_total int64) (err error)
	LogPerMonth(user_id, voucher_id, action string, check_total int64) (err error)
	LogPerYear(user_id, voucher_id, action string, check_total int64) (err error)
	LogAll(user_id, voucher_id, action string, check_total int64) (err error)
}
