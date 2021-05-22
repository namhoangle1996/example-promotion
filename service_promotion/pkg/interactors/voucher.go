package interactors

import (
	"fmt"
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

func (i InteractorImpl) CheckUserCanUseVoucher(user_id, voucher_code string) (*entities.Voucher, *entities.PromotionUser, error) {
	voucher_detail, err := i.IVoucher.FindByCode(voucher_code)

	if err != nil {
		return nil, nil, err
	}

	if voucher_detail == nil {
		return nil, nil, fmt.Errorf("[CheckUserCanUseVoucher] cannot find voucher")
	}

	voucher_of_user, err := i.IUserHasVoucher.FindById(user_id, voucher_detail.Id)

	if err != nil {
		return nil, nil, err
	}

	if voucher_of_user == nil {
		return nil, nil, fmt.Errorf("[CheckUserCanUseVoucher] cannot find user_voucher")
	}

	if voucher_detail.AllowAll == false {
		check, err := i.IUserApplyVoucher.CheckApplied(user_id, voucher_detail.Id)

		if err != nil {
			return nil, nil, err
		}

		if check == 0 {
			return nil, nil, fmt.Errorf("[CheckUserCanUseVoucher] not apply for this user")
		}
	}

	if voucher_of_user.Total == 0 {
		return nil, nil, fmt.Errorf("[CheckUserCanUseVoucher] Voucher user is not enough! ")
	}

	if voucher_detail.ExpiredAt < int64(helpers.GetCurrentTime().Unix()) {
		return nil, nil, fmt.Errorf("[CheckUserCanUseVoucher] Voucher is exprired! ")
	}

	if voucher_detail.StartedAt > int64(helpers.GetCurrentTime().Unix()) {
		return nil, nil, fmt.Errorf("[CheckUserCanUseVoucher] Voucher is not started yet! ")
	}

	return voucher_detail, voucher_of_user, nil
}
