package errors

import err_micro "github.com/micro/go-micro/v2/errors"

const (
	WRONG_AMOUNT = 32
)

func NewErrorMsg(internal string, code int32) error {
	return &err_micro.Error{Detail: internal, Code: code}
}

func NotFoundWallet() error {
	return &err_micro.Error{Detail: "Tài khoản không đủ ", Code: 6}
}

func EnoughTotalInventory() error {
	return &err_micro.Error{Detail: "Đã hết mã voucher ", Code: 20}
}

func VoucherNotExits() error {
	return &err_micro.Error{Detail: "Voucher không tồn tại hoặc đã hết hạn ", Code: 21}
}

func ReverseError() error {
	return &err_micro.Error{Detail: "Tài khoản không đủ ", Code: 23}
}
