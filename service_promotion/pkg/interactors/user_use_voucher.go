package interactors

import (
	"context"
	"fmt"
	"gitlab.com/wallet-gpay/service-promotion/constant"
	"gitlab.com/wallet-gpay/service-promotion/errors"

	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/helpers/saga"
	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

func (i InteractorImpl) UseVoucher(in *service_promotion.UseVoucherRequest, out *service_promotion.UseVoucherResponse) (voucher *entities.Voucher, err error) {
	sg := saga.NewSaga("UseVoucher")

	voucher, _, err = i.CheckUserCanUseVoucher(in.UserId, in.Code)
	if err != nil {
		return
	}

	sg.AddStep(&saga.Step{
		Name: "UseVoucher",
		Func: func(ctx context.Context) (err error) {
			discountAmount := int64(0)
			if voucher.Type != in.ServiceCode {
				err = fmt.Errorf("Invalid service code ")
				return
			}

			if !voucher.SourceOfFunds.IsContains(in.SourceOfFund) {
				err = fmt.Errorf("Invalid source of fund code ")
				return
			}

			if in.SourceOfFund == constant.SOURCE_OF_FUND_BALANCE_WALLET && in.CurrentBalance < in.Amount {
				return fmt.Errorf(errors.NotFoundWallet().Error())
			}

			if in.Amount < voucher.MinTransValue {
				err = fmt.Errorf("Giá trị giao dịch tối thiểu để áp dụng mã khuyến mãi là %v", voucher.MinTransValue)
				return
			}

			err = i.debitWalletVoucher(voucher.Id, in.UserId, in.TraceId, in.Total)

			if voucher.DiscountAmount > 0 {
				discountAmount = voucher.DiscountAmount

			}

			if voucher.Percent > 0 {
				discountAmount = int64(voucher.Percent*float64(in.Amount)) / 100

			}
			if discountAmount > voucher.MaxAmount && voucher.MaxAmount > 0 {
				discountAmount = voucher.MaxAmount
			}
			if discountAmount > in.Amount {
				discountAmount = in.Amount
			}

			out.DiscountAmount = discountAmount
			out.LastAmount = in.Amount - discountAmount
			return
		},
		CompensateFunc: func(c context.Context) (err error) {
			return i.ReverseWallet(in.TraceId)
		},
		Options: nil,
	})

	sg.AddStep(&saga.Step{
		Name: "CheckLogUseVoucher",
		Func: func(c context.Context) (err error) {
			switch voucher.TotalCanUsePer {
			case "day":
				err = i.ILogVoucher.LogPerDay(in.UserId, voucher.Id, "u", voucher.TotalCanUse)
			case "week":
				err = i.ILogVoucher.LogPerWeek(in.UserId, voucher.Id, "u", voucher.TotalCanUse)
			case "month":
				err = i.ILogVoucher.LogPerMonth(in.UserId, voucher.Id, "u", voucher.TotalCanUse)
			case "year":
				err = i.ILogVoucher.LogPerYear(in.UserId, voucher.Id, "u", voucher.TotalCanUse)
			case "all":
				err = i.ILogVoucher.LogAll(in.UserId, voucher.Id, "u", voucher.TotalCanUse)
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

	err = rg.ExecutionError
	return
}
