package log_voucher

import (
	"context"
	"fmt"
	"gitlab.com/wallet-gpay/service-promotion/errors"
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type repoImpl struct {
	collection     *mongo.Collection
	collection_cmd *mongo.Collection
	lockCountD     *sync.Mutex
	lockCountM     *sync.Mutex
	lockCountW     *sync.Mutex
	lockCountA     *sync.Mutex
}

func (r repoImpl) Create(voucher entities.LogUserVoucher) error {
	_, err := r.collection.InsertOne(helpers.ContextWithTimeOut(), voucher)
	return err
}

func (r repoImpl) Update(voucher entities.LogUserVoucher) error {
	_, err := r.collection.InsertOne(helpers.ContextWithTimeOut(), voucher)
	return err
}

func (r repoImpl) CountLogPerDay(user_id, voucher_id, action string) int64 {
	current_time, _ := time.Parse("2006-01-02", helpers.GetCurrentTime().String())
	count, _ := r.collection.CountDocuments(helpers.ContextWithTimeOut(), bson.D{{
		"id_user", user_id,
	}, {"action", action}, {"state", "FINISH"}, {"id_voucher", voucher_id}, {
		"created_at", bson.D{{"$gte", current_time}},
	}})
	return count
}

func (r repoImpl) LogPerDay(user_id, voucher_id, action string, check_total int64) (err error) {
	date := helpers.GetDayOfYear()

	r.lockCountD.Lock()
	defer r.lockCountD.Unlock()
	c, err := r.collection_cmd.CountDocuments(helpers.ContextWithTimeOut(), bson.D{{
		"d", date,
	}, {"user_id", user_id}, {"action", action}, {"voucher_id", voucher_id}})

	upsert := c == 0

	update_res, err := r.collection_cmd.UpdateOne(helpers.ContextWithTimeOut(), bson.M{"$where": fmt.Sprintf("this.user_id == `%v` && this.d == %v && this.action == `%v` && this.total < %v", user_id, date, action, check_total)},
		bson.D{{
			"$inc", bson.M{"total": 1},
		}, {"$set", bson.D{
			{"user_id", user_id},
			{"voucher_id", voucher_id},
			{"action", action},
			{"d", date},
		}}}, &options.UpdateOptions{
			Upsert: &upsert,
		})
	if (update_res == nil || update_res.ModifiedCount == 0) && upsert == false {
		return fmt.Errorf("Quá giới hạn voucher được " + helpers.VoucherActionConvert(action) + " trong ngày")
	}

	return
}

func (r repoImpl) LogPerWeek(user_id, voucher_id, action string, check_total int64) (err error) {
	_, week := helpers.GetCurrentTime().ISOWeek()
	r.lockCountW.Lock()
	defer r.lockCountW.Unlock()
	c, err := r.collection_cmd.CountDocuments(helpers.ContextWithTimeOut(), bson.D{{
		"w", week,
	}, {"user_id", user_id}, {"action", action}, {"voucher_id", voucher_id}})

	upsert := c == 0
	update_res, err := r.collection_cmd.UpdateOne(helpers.ContextWithTimeOut(), bson.M{
		"$where": fmt.Sprintf("this.user_id == `%v` && this.w == %v && this.action == `%v` && this.total < %v", user_id, week, action, check_total),
	}, bson.D{
		{"$set", bson.D{
			{"user_id", user_id},
			{"voucher_id", voucher_id},
			{"action", action},
			{"w", week},
		}},
		{
			"$inc", bson.M{"total": 1},
		},
	}, &options.UpdateOptions{
		Upsert: &upsert,
	})
	if (update_res == nil || update_res.ModifiedCount == 0) && upsert == false {
		return fmt.Errorf("Quá giới hạn voucher được " + helpers.VoucherActionConvert(action) + " trong tuần")
	}
	return
}

func (r repoImpl) LogPerMonth(user_id, voucher_id, action string, check_total int64) (err error) {
	month := int(helpers.GetCurrentTime().Month())
	r.lockCountM.Lock()
	defer r.lockCountM.Unlock()
	c, err := r.collection_cmd.CountDocuments(helpers.ContextWithTimeOut(), bson.D{{
		"m", month,
	}, {"user_id", user_id}, {"action", action}, {"voucher_id", voucher_id}})

	upsert := c == 0

	update_res, err := r.collection_cmd.UpdateOne(helpers.ContextWithTimeOut(), bson.M{
		"$where": fmt.Sprintf("this.user_id ==`%v` && this.m == %v && this.action == `%v` && this.total < %v", user_id, month, action, check_total),
	}, bson.D{
		{"$set", bson.D{
			{"user_id", user_id},
			{"voucher_id", voucher_id},
			{"action", action},
			{"m", month},
		}},
		{
			"$inc", bson.M{"total": 1},
		},
	}, &options.UpdateOptions{
		Upsert: &upsert,
	})
	if (update_res == nil || update_res.ModifiedCount == 0) && upsert == false {
		return fmt.Errorf("Quá giới hạn voucher được " + helpers.VoucherActionConvert(action) + " trong tháng")
	}
	return
}

func (r repoImpl) LogPerYear(user_id, voucher_id, action string, check_total int64) (err error) {
	year, _ := helpers.GetCurrentTime().ISOWeek()
	r.lockCountW.Lock()
	defer r.lockCountW.Unlock()
	update_res, err := r.collection_cmd.UpdateOne(helpers.ContextWithTimeOut(), bson.M{
		"$where": fmt.Sprintf("this.y == %v && this.action == `%v` && this.total < %v", year, action, check_total),
	}, bson.D{
		{"$set", bson.D{
			{"user_id", user_id},
			{"voucher_id", voucher_id},
			{"action", action},
			{"y", year},
		}},
		{
			"$inc", bson.M{"total": 1},
		},
	}, &options.UpdateOptions{})
	if update_res == nil || update_res.ModifiedCount == 0 {
		return fmt.Errorf("CheckLogPerWeek ")
	}
	return
}

func (r repoImpl) LogAll(user_id, voucher_id, action string, check_total int64) (err error) {
	r.lockCountA.Lock()
	c, err := r.collection_cmd.CountDocuments(helpers.ContextWithTimeOut(), bson.D{
		{
			"a", "all",
		}, {"user_id", user_id},
		{"action", action},
		{"voucher_id", voucher_id},
	},
	)

	upsert := c == 0
	defer r.lockCountA.Unlock()

	update_res, err := r.collection_cmd.UpdateOne(helpers.ContextWithTimeOut(), bson.M{
		"$where": fmt.Sprintf("this.user_id == `%v` && this.a == `%v` && this.action == `%v` && this.total < %v && this.voucher_id == `%v`", user_id, "all", action, check_total, voucher_id),
	}, bson.D{
		{"$set", bson.D{
			{"user_id", user_id},
			{"voucher_id", voucher_id},
			{"action", action},
			{"a", "all"},
		}},
		{
			"$inc", bson.M{"total": 1},
		},
	}, &options.UpdateOptions{
		Upsert: &upsert,
	})
	if (update_res == nil || update_res.ModifiedCount == 0) && upsert == false {
		return fmt.Errorf("Quá giới hạn voucher được " + helpers.VoucherActionConvert(action) + " trong chương trình lần này")
	}
	return
}

func (r repoImpl) CountLogAll(user_id, voucher_id, action string) int64 {
	count, _ := r.collection.CountDocuments(helpers.ContextWithTimeOut(), bson.D{{
		"id_user", user_id,
	}, {"action", action}, {"state", "FINISH"}, {"id_voucher", voucher_id}})
	return count
}

func (r repoImpl) FindLogBy(trace_id string, check_reverse bool) (voucher []*entities.LogUserVoucher, err error) {
	ctx := helpers.ContextWithTimeOut()

	if check_reverse == true {
		counter, _ := r.collection.CountDocuments(context.TODO(), bson.D{
			{
				"trace_id", trace_id,
			},
			{
				"action", "REVERSE",
			},
		})

		if counter >= 1 {
			return nil, errors.ReverseError()
		}
	}

	cur, err := r.collection.Find(helpers.ContextWithTimeOut(), bson.D{
		{"trace_id", trace_id},
	})

	if err != nil {
		return
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var voucher_detail *entities.LogUserVoucher
		err_decode := cur.Decode(&voucher_detail)
		if err_decode == nil {
			voucher = append(voucher, voucher_detail)
		}
	}
	return
}

func NewRepository(db *mongo.Client) *repoImpl {
	collection_cmd := db.Database("service_promotion").Collection("promo_log_cmd")
	collection_action := db.Database("service_promotion").Collection("promo_log_user_voucher")

	//index

	//
	return &repoImpl{
		collection:     collection_action,
		collection_cmd: collection_cmd,
		lockCountD:     &sync.Mutex{},
		lockCountM:     &sync.Mutex{},
		lockCountW:     &sync.Mutex{},
		lockCountA:     &sync.Mutex{},
	}
}
