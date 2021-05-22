package mongodb

import "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"

type IPopup interface {
	GetList(user_id string) ([]*entities.PopUp, error)
}
