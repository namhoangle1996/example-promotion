package popup

import (
	"context"
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type repoImpl struct {
	collection  *mongo.Collection
	collection2 *mongo.Collection
}

func (r repoImpl) GetList(user_id string) ([]*entities.PopUp, error) {
	time_now := time.Now().Unix()
	cur, err := r.collection.Find(helpers.ContextWithTimeOut(), bson.D{
		{"start_date", bson.D{{"$lte", time_now}}},
		{"end_date", bson.D{{"$gte", time_now}}},
		{"deleted_at", nil},
	}, &options.FindOptions{
		Sort: bson.D{{"created_at", -1}},
	})

	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var data_popup []*entities.PopUp
	for cur.Next(context.TODO()) {
		var detail_popup *entities.PopUp
		err = cur.Decode(&detail_popup)

		if err == nil {
			if detail_popup.AllowAll == false {
				check, err := r.collection2.CountDocuments(helpers.ContextWithTimeOut(), bson.D{
					{"user_id", user_id},
					{"popup_id", detail_popup.Id},
					{"deleted_at", nil},
				})
				if err == nil && check > 0 {
					data_popup = append(data_popup, detail_popup)
				}
			} else {
				data_popup = append(data_popup, detail_popup)
			}
		}
	}
	return data_popup, nil
}

func NewRepository(db *mongo.Client) *repoImpl {
	return &repoImpl{
		collection:  db.Database("service_promotion").Collection("popups"),
		collection2: db.Database("service_promotion").Collection("popup_user"),
	}
}
