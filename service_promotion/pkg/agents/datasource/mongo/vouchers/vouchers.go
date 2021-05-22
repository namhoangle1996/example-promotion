package vouchers

import (
	"context"
	"fmt"
	"gitlab.com/wallet-gpay/service-promotion/constant"
	"gitlab.com/wallet-gpay/service-promotion/errors"
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repoImpl struct {
	collection *mongo.Collection
}

func (r repoImpl) IncrementTotal(voucher_id string, total uint64) error {
	_, err := r.collection.UpdateOne(helpers.ContextWithTimeOut(), bson.D{
		{"_id", voucher_id},
	}, bson.D{{"$inc", bson.M{"total": total}}})
	return err
}

func (r repoImpl) DecrementTotal(voucher_id string, total uint64) error {
	update, err := r.collection.UpdateOne(helpers.ContextWithTimeOut(), bson.D{
		{"_id", voucher_id},
		{
			"$where", fmt.Sprintf("this.total >= %v", total),
		},
	}, bson.D{
		{
			"$inc", bson.M{"total": -int64(total)},
		},
	},
	)
	if update == nil || update.ModifiedCount == 0 {
		return errors.EnoughTotalInventory()
	}
	return err
}

func (r repoImpl) FindById(id string) (*entities.Voucher, error) {
	var voucher *entities.Voucher
	err := r.collection.FindOne(helpers.ContextWithTimeOut(), bson.D{{"_id", id}, {"status", "ACTIVE"}, {"deleted_at", nil}}).Decode(&voucher)

	if err != nil {
		return nil, errors.VoucherNotExits()
	}
	return voucher, err
}

func (r repoImpl) FindByCode(code string) (*entities.Voucher, error) {
	var voucher *entities.Voucher
	err := r.collection.FindOne(helpers.ContextWithTimeOut(), bson.D{
		{"voucher_code", code},
		{"started_at", bson.M{"$lte": helpers.GetCurrentTime().Unix()}},
		{"expired_at", bson.M{"$gte": helpers.GetCurrentTime().Unix()}},
		{"status", "ACTIVE"},
		{"deleted_at", nil}}).Decode(&voucher)

	if err != nil {
		return nil, errors.VoucherNotExits()
	}

	return voucher, err
}

func (r repoImpl) GetList(in *service_promotion.GetListPromotionRequest) ([]*entities.Voucher, error) {
	var vouchers []*entities.Voucher
	limit := int64(50)
	filter := bson.D{
		{"status", "ACTIVE"},
		{"started_at", bson.M{"$lte": helpers.GetCurrentTime().Unix()}},
		{"expired_at", bson.M{"$gte": helpers.GetCurrentTime().Unix()}},
		{"deleted_at", nil},
	}
	if in.IdGapo == "" { //case not GAPO user
		filter = append(filter, bson.E{Key: "source_user", Value: bson.M{"$ne": constant.GAPO_USER}})
	}
	cur, err := r.collection.Find(helpers.ContextWithTimeOut(), filter,
		&options.FindOptions{Sort: bson.M{"started_at": -1}, Skip: &in.Offset, Limit: &limit},
	)
	if err != nil {
		return nil, err
	}

	if cur == nil {
		return nil, fmt.Errorf("[GetList] Cur nil ")
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var voucer_detail *entities.Voucher
		err = cur.Decode(&voucer_detail)
		if err == nil {
			vouchers = append(vouchers, voucer_detail)
		}
	}
	return vouchers, nil

}

func NewRepository(db *mongo.Client) *repoImpl {
	return &repoImpl{
		collection: db.Database("service_promotion").Collection("promo_voucher"),
	}
}
