package grpc

import (
	"context"

	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (presenter *Service) GetListPopup(
	ctx context.Context,
	in *service_promotion.GetListPopupRequest,
	out *service_promotion.GetListPopupResponse,
) error {
	var service_popups []*service_promotion.PopupDTO
	popups, err := presenter.user_interactor.GetPopupPromtion(in.UserId)

	if err != nil {
		return err
	}

	for _, popup := range popups {
		service_popups = append(service_popups, &service_promotion.PopupDTO{
			Id:          popup.Id,
			Type:        popup.Type,
			Title:       popup.Title,
			Description: popup.Description,
			ImageURL:    popup.ImageURL,
			StartDate:   int64(popup.StartDate),
			EndDate:     int64(popup.EndDate),
			Frequency:   int64(popup.Frequency),
			ExternalURL: popup.ExternalURL,
			AllowAll:    popup.AllowAll,
		})
	}
	out.Popups = service_popups
	return nil

}
