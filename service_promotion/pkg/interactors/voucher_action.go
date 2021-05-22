package interactors

import (
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

func (i InteractorImpl) debitWalletVoucher(voucher_id, user_id, trace_id string, total int64) (err error) {
	err = i.IUserHasVoucher.DecrementTotalVoucher(voucher_id, user_id, total)

	if err != nil {
		return
	}

	err = i.ILogVoucher.Create(entities.LogUserVoucher{
		Id:        helpers.BusinessID("LOG"),
		IdVoucher: voucher_id,
		IdUser:    user_id,
		Total:     total,
		Action:    "DEBIT",
		TraceID:   trace_id,
		CreatedAt: helpers.GetCurrentTime(),
		UpdatedAt: helpers.GetCurrentTime(),
	})

	return
}

func (i InteractorImpl) creditWalletVoucher(
	voucher_id,
	user_id,
	trace_id string,
	total int64) (err error) {
	promotion_user_has_voucher, err := i.IUserHasVoucher.FindById(user_id, voucher_id)

	// Tạo ví voucher
	if promotion_user_has_voucher != nil {
		err = i.IUserHasVoucher.IncrementTotalVoucher(voucher_id, user_id, total)
	} else {
		err = i.IUserHasVoucher.Create(entities.PromotionUser{
			Id:        helpers.BusinessID("PRU"),
			IdVoucher: voucher_id,
			IdUser:    user_id,
			Total:     total,
			CreatedAt: helpers.GetCurrentTime(),
			LastBuy:   helpers.GetCurrentTime(),
			UpdatedAt: helpers.GetCurrentTime(),
		})
	}
	if err != nil {
		return
	}

	err = i.ILogVoucher.Create(entities.LogUserVoucher{
		Id:        helpers.BusinessID("LOG"),
		IdVoucher: voucher_id,
		IdUser:    user_id,
		Total:     total,
		Action:    "CREDIT",
		TraceID:   trace_id,
		CreatedAt: helpers.GetCurrentTime(),
		UpdatedAt: helpers.GetCurrentTime(),
	})
	return
}

func (i InteractorImpl) debitGPAY(voucher_id string, trace_id string, total uint64) (err error) {
	err = i.IVoucher.DecrementTotal(voucher_id, total)

	if err != nil {
		return
	}

	err = i.ILogVoucher.Create(entities.LogUserVoucher{
		Id:        helpers.BusinessID("LOG"),
		IdVoucher: voucher_id,
		Total:     int64(total),
		GpayId:    "GPAY0001",
		Action:    "DEBIT",
		TraceID:   trace_id,
		CreatedAt: helpers.GetCurrentTime(),
		UpdatedAt: helpers.GetCurrentTime(),
	})
	return
}

func (i InteractorImpl) creditGPAY(voucher_id, trace_id string, total uint64) (err error) {
	err = i.IVoucher.IncrementTotal(voucher_id, total)

	if err != nil {
		return
	}

	err = i.ILogVoucher.Create(entities.LogUserVoucher{
		Id:        helpers.BusinessID("LOG"),
		IdVoucher: voucher_id,
		Total:     int64(total),
		GpayId:    "GPAY0001",
		Action:    "CREDIT",
		TraceID:   trace_id,
		CreatedAt: helpers.GetCurrentTime(),
		UpdatedAt: helpers.GetCurrentTime(),
	})

	return
}

func (i InteractorImpl) ReverseWallet(trace_id string) (err error) {
	logs, err := i.ILogVoucher.FindLogBy(trace_id , true )

	for k, log := range logs {
		if log.Action == "CREDIT" {
			if log.GpayId != "" {
				err = i.debitGPAY(log.IdVoucher, trace_id, uint64(log.Total))
			}

			if log.IdUser != "" {
				err = i.debitWalletVoucher(log.IdVoucher, log.IdUser, trace_id, log.Total)
			}

		} else if log.Action == "DEBIT" {
			if log.GpayId != "" {
				err = i.creditGPAY(log.IdVoucher, trace_id, uint64(log.Total))
			}

			if log.IdUser != "" {
				err = i.creditWalletVoucher(log.IdVoucher, log.IdUser, trace_id, log.Total)
			}
		}
		if k == len(logs)-1 {
			//Log REVERSE
			err = i.ILogVoucher.Create(entities.LogUserVoucher{
				Id:        helpers.BusinessID("LOG"),
				IdVoucher: log.IdVoucher,
				Action:    "REVERSE",
				TraceID:   trace_id,
				CreatedAt: helpers.GetCurrentTime(),
				UpdatedAt: helpers.GetCurrentTime(),
			})
		}

	}

	return
}
