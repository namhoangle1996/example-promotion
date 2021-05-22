package grpc

import (
	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

func ConvertEntitiesVoucher2DTO(entities_voucher *entities.Voucher) *service_promotion.VoucherDTO {

	return &service_promotion.VoucherDTO{
		Id:                 entities_voucher.Id,
		Title:              entities_voucher.Title,
		Thumbnail:          entities_voucher.Thumbnail,
		Rule:               entities_voucher.Rule,
		Description:        entities_voucher.Description,
		ShortDescription:   entities_voucher.ShortDescription,
		Code:               entities_voucher.Code,
		Amount:             int64(entities_voucher.Amount),
		Type:               entities_voucher.Type,
		Total:              entities_voucher.Total,
		TotalFreeze:        entities_voucher.TotalFreeze,
		FreezeIds:          entities_voucher.FreezeIds,
		TotalCanBuy:        entities_voucher.TotalCanBuy,
		TotalCanBuyPerDay:  entities_voucher.TotalCanBuyPerDay,
		TotalCanUsePerDay:  entities_voucher.TotalCanUsePerDay,
		TotalCanUsePerWeek: entities_voucher.TotalCanUsePerWeek,
		TotalCanUse:        entities_voucher.TotalCanUse,
		DiscountAmount:     uint64(entities_voucher.DiscountAmount),
		DiscountPercent:    entities_voucher.Percent,
		MerchantID:         entities_voucher.MerchantID,
		MaxAmount:          uint64(entities_voucher.MaxAmount),
		Status:             entities_voucher.Status,
		Show:               entities_voucher.Show,
		AllowAll:           entities_voucher.AllowAll,
		ShowInMyVoucher:    entities_voucher.ShowInMyVoucher,
		StartedAt:          entities_voucher.StartedAt,
		ExpiredAt:          entities_voucher.ExpiredAt,
		MinTransValue:      entities_voucher.MinTransValue,
	}
}

func ConvertEntitiesMerchant2DTO(entities *entities.Merchant) *service_promotion.MerchantDTO {
	return &service_promotion.MerchantDTO{
		Id:          entities.Id,
		Email:       entities.Email,
		Code:        entities.Code,
		Logo:        entities.Logo,
		Address:     entities.Address,
		Name:        entities.Name,
		PhoneNumber: entities.PhoneNumber,
		Website:     entities.Website,
	}
}
