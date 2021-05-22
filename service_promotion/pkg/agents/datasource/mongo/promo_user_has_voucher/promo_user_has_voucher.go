package promo_user_has_voucher

import (
	"fmt"
	"gitlab.com/wallet-gpay/service-promotion/errors"
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/helpers/mongoindex"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type repoImpl struct {
	collection *mongo.Collection
}

func (r repoImpl) Get(user_id string, offset int64) ([]*entities.PromotionUser, error) {
	var vouchers_of_user []*entities.PromotionUser
	limit := int64(10)
	cur, err := r.collection.Find(helpers.ContextWithTimeOut(), bson.D{
		{"id_user", user_id}},
		&options.FindOptions{
			Sort:  bson.M{"updated_at": -1},
			Skip:  &offset,
			Limit: &limit,
		},
	)

	if err != nil {
		fmt.Println("errrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr", err)
		return nil, err
	}
	defer cur.Close(helpers.ContextWithTimeOut())
	for cur.Next(helpers.ContextWithTimeOut()) {
		var voucher_of_user *entities.PromotionUser
		err = cur.Decode(&voucher_of_user)
		if err != nil {
			fmt.Println("err got", err)
		}
		if err == nil {
			vouchers_of_user = append(vouchers_of_user, voucher_of_user)
		}
	}

	return vouchers_of_user, nil
}

func (r repoImpl) FindByCode(user_id, code string) (*entities.PromotionUser, error) {
	var voucher_of_user *entities.PromotionUser

	err := r.collection.FindOne(helpers.ContextWithTimeOut(), bson.D{
		{
			"code", code,
		},
		{"id_user", user_id}},
		&options.FindOneOptions{
			Sort: bson.M{"created_at": -1},
		},
	).Decode(&voucher_of_user)
	return voucher_of_user, err
}

func (r repoImpl) FindById(user_id, id string) (*entities.PromotionUser, error) {
	var voucher_of_user *entities.PromotionUser

	err := r.collection.FindOne(helpers.ContextWithTimeOut(), bson.D{
		{
			"id_voucher", id,
		},
		{"id_user", user_id}},
		&options.FindOneOptions{
			Sort: bson.M{"created_at": -1},
		},
	).Decode(&voucher_of_user)

	return voucher_of_user, err
}

func (r repoImpl) Create(user_voucher entities.PromotionUser) error {
	_, err := r.collection.InsertOne(helpers.ContextWithTimeOut(), user_voucher)
	return err
}

func (r repoImpl) Update(user_voucher entities.PromotionUser) error {
	panic("implement me")
}

func (r repoImpl) DecrementTotalVoucher(id, user_id string, total int64) error {
	update, err := r.collection.UpdateOne(helpers.ContextWithTimeOut(), bson.D{
		{"$where", fmt.Sprintf("this.id_user == `%v` && this.total >= %v & this.id_voucher == `%v`", user_id, total, id)},
	},

		bson.M{
			"$inc": bson.M{"total": -total},
		},
	)
	if update == nil || update.ModifiedCount == 0 {
		return errors.EnoughTotalInventory()
	}
	return err
}

func (r repoImpl) IncrementTotalVoucher(id, user_id string, total int64) error {
	_, err := r.collection.UpdateOne(helpers.ContextWithTimeOut(), bson.D{
		{"$where", fmt.Sprintf("this.id_user == `%v` &&this.id_voucher == `%v`", user_id, id)},
	},
		bson.M{
			"$inc": bson.M{"total": total},
			"$set": bson.M{"updated_at": time.Now()},
		},
	)
	return err
}

func NewRepository(db *mongo.Client) *repoImpl {
	c := db.Database("service_promotion").Collection("promo_user_voucher")
	ctx := helpers.ContextWithTimeOut()
	mongoindex.EnsureIndex(ctx, c, []bson.E{
		{Key: "status", Value: -1},
	}, false)

	mongoindex.EnsureIndex(ctx, c, []bson.E{
		{Key: "id_user", Value: -1},
	}, false)

	mongoindex.EnsureIndex(ctx, c, []bson.E{
		{Key: "created_at", Value: -1},
	}, false)

	mongoindex.EnsureIndex(ctx, c, []bson.E{
		{Key: "updated_at", Value: -1},
	}, false)

	mongoindex.EnsureIndex(ctx, c, []bson.E{
		{Key: "code", Value: 1},
		{Key: "status", Value: 1},
		{Key: "id_user", Value: 1},
	}, false)

	return &repoImpl{
		collection: c,
	}
}
