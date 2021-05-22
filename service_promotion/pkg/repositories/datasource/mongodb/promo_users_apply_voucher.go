package mongodb

type IUserApplyVoucher interface {
	CheckApplied(userid, promoid string) (int64, error)
}
