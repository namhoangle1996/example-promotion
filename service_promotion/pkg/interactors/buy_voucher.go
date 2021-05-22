package interactors

import (
	"context"
	"fmt"
	"gitlab.com/wallet-gpay/service-promotion/constant"
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/helpers/saga"
	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
)

func (i InteractorImpl) BuyVoucher(in *service_promotion.BuyVoucherRequest) error {
	sg := saga.NewSaga("ActionAccountingPromotionDeposit")
	voucher_detail, err := i.IVoucher.FindById(in.VoucherId)

	if err != nil {
		return err
	}

	if voucher_detail == nil {
		return fmt.Errorf("[BuyVoucher] cannot find voucher buy")
	}

	if voucher_detail.SourceUser == constant.GAPO_USER && in.IdGapo == "" {
		return fmt.Errorf("Voucher chỉ dành cho người dùng GAPO")
	}

	if voucher_detail.AllowAll == false {

		check, _ := i.IUserApplyVoucher.CheckApplied(in.UserId, in.VoucherId)

		if check == 0 {
			return fmt.Errorf("[BuyVoucher] voucher cannot buy by this user")
		}

	}

	sg.AddStep(&saga.Step{
		Name: "DECREMENT_VOUCHER",
		Func: func(c context.Context) (err error) {
			err = i.debitGPAY(in.VoucherId, in.TraceId, uint64(in.Total))
			return
		},
		CompensateFunc: func(c context.Context) (err error) {
			return i.ReverseWallet(in.TraceId)
		},
		Options: nil,
	})

	sg.AddStep(&saga.Step{
		Name: "INCREMENT_USER_VOUCHER",
		Func: func(c context.Context) (err error) {
			err = i.creditWalletVoucher(in.VoucherId, in.UserId, in.TraceId, in.Total)
			return
		},
		CompensateFunc: func(c context.Context) (err error) {
			return
		},
		Options: nil,
	})

	sg.AddStep(&saga.Step{
		Name: "CheckLogBuy",
		Func: func(c context.Context) (err error) {
			if voucher_detail.TotalCanBuy > 0 {
				switch voucher_detail.TotalCanBuyPer {
				case "day":
					err = i.ILogVoucher.LogPerDay(in.UserId, voucher_detail.Id, "b", voucher_detail.TotalCanBuy)
				case "week":
					err = i.ILogVoucher.LogPerWeek(in.UserId, voucher_detail.Id, "b", voucher_detail.TotalCanBuy)
				case "month":
					err = i.ILogVoucher.LogPerMonth(in.UserId, voucher_detail.Id, "b", voucher_detail.TotalCanBuy)
				case "year":
					err = i.ILogVoucher.LogPerYear(in.UserId, voucher_detail.Id, "b", voucher_detail.TotalCanBuy)
				case "all":
					err = i.ILogVoucher.LogAll(in.UserId, voucher_detail.Id, "b", voucher_detail.TotalCanBuy)
				}
			}

			return
		},
		CompensateFunc: func(c context.Context) (err error) {
			return
		},
		Options: nil,
	})

	ordinator := saga.NewCoordinator(helpers.ContextWithTimeOut(), helpers.ContextWithTimeOut(), sg, i.LogSaga)
	rg := ordinator.Play()

	return rg.ExecutionError
}
