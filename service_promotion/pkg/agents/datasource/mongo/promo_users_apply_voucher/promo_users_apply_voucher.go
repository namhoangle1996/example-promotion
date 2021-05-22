package promo_users_apply_voucher

import (
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type repoImpl struct {
	collection *mongo.Collection
}

func (r repoImpl) CheckApplied(userid, promoid string) (int64, error) {
	var apply_users *entities.PromotionUserApply
	c := int64(0)
	//Check time
	err := r.collection.FindOne(helpers.ContextWithTimeOut(), bson.D{
		{"promo_id", promoid},
		{"user_id", userid},
	}).Decode(&apply_users)

	if err != nil {
		return 0, err
	}

	if apply_users != nil {
		c = 1
	}

	if apply_users.StartDate != 0 && apply_users.EndDate != 0 {
		if time.Now().Unix() >= apply_users.StartDate && time.Now().Unix() <= apply_users.EndDate {
			c = apply_users.EndDate
		} else if time.Now().Unix() < apply_users.StartDate || time.Now().Unix() > apply_users.EndDate {
			c = 0
		}
	}


	return c, nil
}

func NewRepository(db *mongo.Client) *repoImpl {
	return &repoImpl{
		collection: db.Database("service_promotion").Collection("promo_users"),
	}
}
