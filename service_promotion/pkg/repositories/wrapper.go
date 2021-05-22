package repositories

import (
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/repositories/data_service"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/repositories/datasource/mongodb"
)

type HandlerRepository struct {
	IVoucher          mongodb.IVoucher
	ILogVoucher       mongodb.ILogVoucher
	IUserApplyVoucher mongodb.IUserApplyVoucher
	IUserHasVoucher   mongodb.IUserHasVoucher
	IPopup            mongodb.IPopup
	IServiceUser      data_service.IServiceUser
}

func NewHandlerRepository(IVoucher mongodb.IVoucher, ILogVoucher mongodb.ILogVoucher, IUserApplyVoucher mongodb.IUserApplyVoucher, IUserVoucher mongodb.IUserHasVoucher, IServiceUser data_service.IServiceUser, IPopup mongodb.IPopup) *HandlerRepository {
	return &HandlerRepository{
		IVoucher,
		ILogVoucher,
		IUserApplyVoucher,
		IUserVoucher,
		IPopup,
		IServiceUser,
	}
}
