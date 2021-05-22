package interactors

import "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"

func (i InteractorImpl) GetPopupPromtion(user_id string) ([]*entities.PopUp, error) {
	return i.IPopup.GetList(user_id)
}
