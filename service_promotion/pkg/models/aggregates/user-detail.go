package aggregates

import (
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

type UserDetail struct {
	User     entities.User
	Balances map[int]entities.Balances
}
