package interactors

import (
	"fmt"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/wallet-gpay/service-promotion/configs"
	"gitlab.com/wallet-gpay/service-promotion/constant"
	"gitlab.com/wallet-gpay/service-promotion/errors"
	"gitlab.com/wallet-gpay/service-promotion/gpooling"
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/helpers/saga"
	"gitlab.com/wallet-gpay/service-promotion/logger"
	"gitlab.com/wallet-gpay/service-promotion/proto/service_promotion"
	infrastructureMongo "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo/log_voucher"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo/promo_user_has_voucher"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo/promo_users_apply_voucher"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo/vouchers"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/repositories"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/repositories/data_service/mocks"
	mocks2 "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/repositories/datasource/mongodb/mocks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InteractorTest struct {
	DBPromotion            *mongo.Database
	Interactor             *InteractorImpl
	MockUser               *mocks.IServiceUser
	MockVoucher            *mocks2.IVoucher
	MockLogVoucher         *mocks2.ILogVoucher
	MockPopup              *mocks2.IPopup
	MockUserHasVoucher     *mocks2.IUserHasVoucher
	MockerUserApplyVoucher *mocks2.IUserApplyVoucher
}

func InitTest() InteractorTest {
	pool, err := gpooling.NewPooling(1000)
	config, _ := configs.LoadTestConfig("./../../../")
	primaryDb := infrastructureMongo.NewMongoDBconnection(config.MongoURI)

	if err != nil {
		panic(err.Error())
	}

	voucher_impl := &mocks2.IVoucher{}
	log_voucher_impl := &mocks2.ILogVoucher{}
	popup_impl := &mocks2.IPopup{}
	user_impl := &mocks.IServiceUser{}
	apply_user_impl := &mocks2.IUserApplyVoucher{}
	user_has_voucher := &mocks2.IUserHasVoucher{}
	log, err := logger.NewLogger("", "production")

	th := &InteractorImpl{
		HandlerRepository: &repositories.HandlerRepository{
			IVoucher:          vouchers.NewRepository(primaryDb),
			ILogVoucher:       log_voucher.NewRepository(primaryDb),
			IUserApplyVoucher: promo_users_apply_voucher.NewRepository(primaryDb),
			IUserHasVoucher:   promo_user_has_voucher.NewRepository(primaryDb),
			IPopup:            popup_impl,
			IServiceUser:      user_impl,
		},
		IPool:   pool,
		ILogger: log,
		LogSaga: saga.New(),
	}

	return InteractorTest{
		DBPromotion:            primaryDb.Database("service_promotion"),
		Interactor:             th,
		MockUser:               user_impl,
		MockVoucher:            voucher_impl,
		MockLogVoucher:         log_voucher_impl,
		MockPopup:              popup_impl,
		MockUserHasVoucher:     user_has_voucher,
		MockerUserApplyVoucher: apply_user_impl,
	}
}
func initPromotion(db *mongo.Database) {
	_, err := db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher",
		Code:           "Voucher",
		Status:         "ACTIVE",
		Total:          1000,
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		TotalCanUsePer: "month",
		TotalCanUse:    1,
		Percent:        1,
		SourceOfFunds:  []string{"TEST"},
	})
	fmt.Println(err)

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:            "Voucher1_allow_false",
		Code:          "Voucher",
		Status:        "ACTIVE",
		Total:         1000,
		AllowAll:      false,
		StartedAt:     helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:     helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds: []string{"TEST"},
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:            "Voucher_3_doesn't_start",
		Code:          "Voucher",
		Status:        "ACTIVE",
		Total:         1000,
		AllowAll:      false,
		StartedAt:     helpers.GetCurrentTime().Add(time.Hour).Unix(),
		ExpiredAt:     helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds: []string{"TEST"},
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:            "Voucher_4_not_enough",
		Code:          "Voucher",
		Status:        "ACTIVE",
		Total:         0,
		AllowAll:      false,
		StartedAt:     helpers.GetCurrentTime().Add(time.Hour).Unix(),
		ExpiredAt:     helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds: []string{"TEST"},
	})

	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "PromotionUser",
		IdUser:    "USER",
		IdVoucher: "Voucher",
		Code:      "Voucher",
		Total:     10,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_BUY_PER_ALL",
		Code:           "Voucher_CAN_BUY_PER_ALL",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanBuy:    1,
		TotalCanBuyPer: "all",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{"TEST"},
	})
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "Voucher_CAN_BUY_PER_ALL",
		IdUser:    "Voucher_CAN_BUY_PER_ALL",
		IdVoucher: "Voucher_CAN_BUY_PER_ALL",
		Code:      "Voucher_CAN_BUY_PER_ALL",
		Total:     1,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_BUY_PER_DAY",
		Code:           "Voucher",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanBuy:    3,
		TotalCanBuyPer: "day",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{"TEST"},
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_BUY_PER_DAY_TOTAL_CAN_BUY_1",
		Code:           "Voucher_CAN_BUY_PER_DAY_TOTAL_CAN_BUY_1",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanBuy:    1,
		TotalCanBuyPer: "day",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_USE_PER_WEEK",
		Code:           "Voucher_CAN_USE_PER_WEEK",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		TotalCanUsePer: "week",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{"TEST"},
	})
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "Voucher_CAN_USE_PER_WEEK",
		IdUser:    "Voucher_CAN_USE_PER_WEEK",
		IdVoucher: "Voucher_CAN_USE_PER_WEEK",
		Code:      "Voucher_CAN_USE_PER_WEEK",
		Total:     1,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_USE_PER_MONTH",
		Code:           "Voucher_CAN_USE_PER_MONTH",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		TotalCanUsePer: "month",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{"TEST"},
	})
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "Voucher_CAN_USE_PER_MONTH",
		IdUser:    "Voucher_CAN_USE_PER_MONTH",
		IdVoucher: "Voucher_CAN_USE_PER_MONTH",
		Code:      "Voucher_CAN_USE_PER_MONTH",
		Total:     1,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_USE_PER_DAY",
		Code:           "Voucher_CAN_USE_PER_DAY",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		TotalCanUsePer: "day",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{"TEST"},
	})
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "Voucher_CAN_USE_PER_DAY",
		IdUser:    "Voucher_CAN_USE_PER_DAY",
		IdVoucher: "Voucher_CAN_USE_PER_DAY",
		Code:      "Voucher_CAN_USE_PER_DAY",
		Total:     1,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_USE_PER_ALL",
		Code:           "Voucher_CAN_USE_PER_ALL",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		TotalCanUsePer: "all",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{"TEST"},
	})
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "Voucher_CAN_USE_PER_ALL",
		IdUser:    "Voucher_CAN_USE_PER_ALL",
		IdVoucher: "Voucher_CAN_USE_PER_ALL",
		Code:      "Voucher_CAN_USE_PER_ALL",
		Total:     1,
	})

	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "test_id11",
		IdUser:    "namle11",
		IdVoucher: "voucher12",
		Code:      "VOUCHER_12",
		Total:     1,
	})

	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "test_id",
		IdUser:    "namle12345",
		IdVoucher: "Voucher_CAN_USE_PER_ALL",
		Code:      "namle",
		Total:     1,
	})

	// check use voucher min value transaction
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "test_id11",
		IdUser:    "namle11",
		IdVoucher: "voucher11",
		Code:      "VOUCHER_11",
		Total:     1,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "voucher11",
		Code:           "VOUCHER_11",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		MinTransValue:  10000,
		TotalCanUsePer: "day",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{"TEST"},
	})

	// Check list source GAPO
	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "id1111",
		Code:           "VOUCHER111",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    10,
		TotalCanUsePer: "all",
		AllowAll:       true,
		SourceUser:     constant.GAPO_USER,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{"TEST"},
	})

	// Check Voucher Per all
	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_BUY_ALL",
		Code:           "Voucher_BUY_ALL",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanBuy:    5,
		TotalCanBuyPer: "all",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{"TEST"},
	})

}
func initPromotionWithRealSOF(db *mongo.Database) {
	_, err := db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher",
		Code:           "Voucher",
		Status:         "ACTIVE",
		Total:          1000,
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		TotalCanUsePer: "month",
		TotalCanUse:    1,
		Percent:        1,
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM, constant.SOURCE_OF_FUND_BALANCE_WALLET},
	})
	fmt.Println(err)

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:            "Voucher1_allow_false",
		Code:          "Voucher",
		Status:        "ACTIVE",
		Total:         1000,
		AllowAll:      false,
		StartedAt:     helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:     helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds: []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:            "Voucher_3_doesn't_start",
		Code:          "Voucher",
		Status:        "ACTIVE",
		Total:         1000,
		AllowAll:      false,
		StartedAt:     helpers.GetCurrentTime().Add(time.Hour).Unix(),
		ExpiredAt:     helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds: []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:            "Voucher_4_not_enough",
		Code:          "Voucher",
		Status:        "ACTIVE",
		Total:         0,
		AllowAll:      false,
		StartedAt:     helpers.GetCurrentTime().Add(time.Hour).Unix(),
		ExpiredAt:     helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds: []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})

	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "PromotionUser",
		IdUser:    "USER",
		IdVoucher: "Voucher",
		Code:      "Voucher",
		Total:     10,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_BUY_PER_ALL",
		Code:           "Voucher_CAN_BUY_PER_ALL",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanBuy:    1,
		TotalCanBuyPer: "all",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "Voucher_CAN_BUY_PER_ALL",
		IdUser:    "Voucher_CAN_BUY_PER_ALL",
		IdVoucher: "Voucher_CAN_BUY_PER_ALL",
		Code:      "Voucher_CAN_BUY_PER_ALL",
		Total:     1,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_BUY_PER_DAY",
		Code:           "Voucher",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanBuy:    3,
		TotalCanBuyPer: "day",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_BUY_PER_DAY_TOTAL_CAN_BUY_1",
		Code:           "Voucher_CAN_BUY_PER_DAY_TOTAL_CAN_BUY_1",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanBuy:    1,
		TotalCanBuyPer: "day",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_USE_PER_WEEK",
		Code:           "Voucher_CAN_USE_PER_WEEK",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		TotalCanUsePer: "week",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "Voucher_CAN_USE_PER_WEEK",
		IdUser:    "Voucher_CAN_USE_PER_WEEK",
		IdVoucher: "Voucher_CAN_USE_PER_WEEK",
		Code:      "Voucher_CAN_USE_PER_WEEK",
		Total:     1,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_USE_PER_MONTH",
		Code:           "Voucher_CAN_USE_PER_MONTH",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		TotalCanUsePer: "month",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "Voucher_CAN_USE_PER_MONTH",
		IdUser:    "Voucher_CAN_USE_PER_MONTH",
		IdVoucher: "Voucher_CAN_USE_PER_MONTH",
		Code:      "Voucher_CAN_USE_PER_MONTH",
		Total:     1,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_USE_PER_DAY",
		Code:           "Voucher_CAN_USE_PER_DAY",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		TotalCanUsePer: "day",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "Voucher_CAN_USE_PER_DAY",
		IdUser:    "Voucher_CAN_USE_PER_DAY",
		IdVoucher: "Voucher_CAN_USE_PER_DAY",
		Code:      "Voucher_CAN_USE_PER_DAY",
		Total:     1,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_CAN_USE_PER_ALL",
		Code:           "Voucher_CAN_USE_PER_ALL",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		TotalCanUsePer: "all",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "Voucher_CAN_USE_PER_ALL",
		IdUser:    "Voucher_CAN_USE_PER_ALL",
		IdVoucher: "Voucher_CAN_USE_PER_ALL",
		Code:      "Voucher_CAN_USE_PER_ALL",
		Total:     1,
	})

	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "test_id",
		IdUser:    "namle12345",
		IdVoucher: "Voucher_CAN_USE_PER_ALL",
		Code:      "namle",
		Total:     1,
	})

	// check use voucher min value transaction
	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "test_id11",
		IdUser:    "namle11",
		IdVoucher: "voucher11",
		Code:      "VOUCHER_11",
		Total:     1,
	})

	db.Collection("promo_user_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.PromotionUser{
		Id:        "test_id12",
		IdUser:    "namle11",
		IdVoucher: "voucher12",
		Code:      "VOUCHER_12",
		Total:     1,
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "voucher11",
		Code:           "VOUCHER_11",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		MinTransValue:  10000,
		TotalCanUsePer: "day",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM, constant.SOURCE_OF_FUND_BALANCE_WALLET},
	})

	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "voucher12",
		Code:           "VOUCHER_12",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    1,
		MinTransValue:  10000,
		TotalCanUsePer: "day",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM, constant.SOURCE_OF_FUND_BALANCE_WALLET},
	})

	// Check list source GAPO
	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "id1111",
		Code:           "VOUCHER111",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanUse:    10,
		TotalCanUsePer: "all",
		AllowAll:       true,
		SourceUser:     constant.GAPO_USER,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})

	// Check Voucher Per all
	db.Collection("promo_voucher").InsertOne(helpers.ContextWithTimeOut(), &entities.Voucher{
		Id:             "Voucher_BUY_ALL",
		Code:           "Voucher_BUY_ALL",
		Status:         "ACTIVE",
		Total:          1000,
		TotalCanBuy:    5,
		TotalCanBuyPer: "all",
		AllowAll:       true,
		StartedAt:      helpers.GetCurrentTime().Add(-time.Hour).Unix(),
		ExpiredAt:      helpers.GetCurrentTime().Add(time.Hour * 10).Unix(),
		SourceOfFunds:  []string{constant.SOURCE_OF_FUND_BANK_ATM},
	})

}

func TestInteractorImpl_CheckUserCanUseVoucher(t *testing.T) {
	th := InitTest()

	initPromotion(th.DBPromotion)

	defer th.DBPromotion.Drop(helpers.ContextWithTimeOut())

	type args struct {
		user_id      string
		voucher_code string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		err      error
		mockFunc func()
	}{
		{
			args: args{
				user_id:      "Voucher_CAN_USE_PER_MONTH",
				voucher_code: "Voucher_CAN_USE_PER_MONTH",
			},
			name:    "CASE CAN USE PER MONTH",
			wantErr: true,
			err:     fmt.Errorf("Quá giới hạn voucher được " + helpers.VoucherActionConvert("u") + " trong tháng"),
			mockFunc: func() {
			},
		},
		{
			args: args{
				user_id:      "Voucher_CAN_USE_PER_WEEK",
				voucher_code: "Voucher_CAN_USE_PER_WEEK",
			},
			name:    "CASE CAN USE PER WEEK",
			wantErr: true,
			err:     fmt.Errorf("Quá giới hạn voucher được " + helpers.VoucherActionConvert("u") + " trong tuần"),
			mockFunc: func() {
			},
		},
		{args: args{
			user_id:      "Voucher_CAN_USE_PER_DAY",
			voucher_code: "Voucher_CAN_USE_PER_DAY",
		},
			name:    "CASE CAN USE PER DAY",
			wantErr: true,
			err:     fmt.Errorf("Quá giới hạn voucher được " + helpers.VoucherActionConvert("u") + " trong ngày"),
			mockFunc: func() {
			},
		},
		{args: args{
			user_id:      "Voucher_CAN_USE_PER_ALL",
			voucher_code: "Voucher_CAN_USE_PER_ALL",
		},
			name:    "CASE CAN USE PER ALL",
			wantErr: true,
			err:     fmt.Errorf("Quá giới hạn voucher được " + helpers.VoucherActionConvert("u") + " trong chương trình lần này"),
			mockFunc: func() {
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			_, _, err := th.Interactor.CheckUserCanUseVoucher(tt.args.user_id, tt.args.voucher_code)
			if err != nil && tt.wantErr == false {
				t.Errorf("error %s", err.Error())
			} else if err != nil && tt.wantErr == true {
				assert.Equal(t, err, tt.err)
			}
		})
	}
}

func TestInteractorImpl_BuyVoucher(t *testing.T) {
	th := InitTest()
	initPromotion(th.DBPromotion)

	defer th.DBPromotion.Drop(helpers.ContextWithTimeOut())

	type args struct {
		user_id    string
		voucher_id string
		total      int64
		trace_id   string
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		err           error
		mockFunc      func()
		expectedCount int64
	}{
		{
			args: args{
				user_id:    "USER",
				voucher_id: "Voucher_CAN_BUY_PER_DAY",
				total:      1,
				trace_id:   "test",
			},
			name: "Voucher_CAN_BUY_PER_DAY",
			mockFunc: func() {
				th.MockUser.On("FindById", mock.AnythingOfType("string")).Return(&entities.User{
					Id: "test1",
				}, nil).Once()
			},
			wantErr: true,
			err:     fmt.Errorf("Quá giới hạn voucher được " + helpers.VoucherActionConvert("b") + " trong ngày"),
		},
		{
			args: args{
				user_id:    "USER",
				voucher_id: "Voucher_CAN_BUY_PER_ALL",
				total:      1,
				trace_id:   "test",
			},
			name: "Voucher_CAN_BUY_PER_ALL",
			mockFunc: func() {
				th.MockUser.On("FindById", mock.AnythingOfType("string")).Return(&entities.User{
					Id: "test1",
				}, nil).Once()
			},
			wantErr: true,
			err:     fmt.Errorf("Quá giới hạn voucher được " + helpers.VoucherActionConvert("b") + " trong chương trình lần này"),
		},
		{
			args: args{
				user_id:    "USER",
				voucher_id: "id1111",
				total:      1,
				trace_id:   "test1111",
			},
			name: "Voucher_CAN_NOT_BUY_FOR_GAPO",
			mockFunc: func() {
				th.MockUser.On("FindById", mock.AnythingOfType("string")).Return(&entities.User{
					Id:     "test1",
					Source: "app",
				}, nil).Once()
			},
			wantErr: true,
			err:     fmt.Errorf("Voucher chỉ dành cho người dùng GAPO"),
		},
		{
			args: args{
				user_id:    "test123456",
				voucher_id: "Voucher_BUY_ALL",
				total:      1,
				trace_id:   "test123456",
			},
			name: "CHECK COUNT BUY PER ALL",
			mockFunc: func() {
				th.MockUser.On("FindById", mock.AnythingOfType("string")).Return(&entities.User{
					Id: "test123456",
				}, nil).Once()
			},
			wantErr:       false,
			expectedCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			err := th.Interactor.BuyVoucher(&service_promotion.BuyVoucherRequest{
				VoucherId: tt.args.voucher_id,
				UserId:    tt.args.user_id,
				Total:     tt.args.total,
				TraceId:   tt.args.trace_id,
			})
			if err != nil && tt.wantErr == false {
				t.Errorf("error %s", err.Error())
			} else if err != nil && tt.wantErr == true {
				assert.Equal(t, err, tt.err)
			} else if err == nil && tt.expectedCount != 0 {
				filter := bson.M{"user_id": tt.args.user_id, "voucher_id": tt.args.voucher_id, "action": "b"}
				gotCountBuyVoucher, _ := th.DBPromotion.Collection("promo_log_cmd").CountDocuments(helpers.ContextWithTimeOut(), filter)
				assert.Equal(t, gotCountBuyVoucher, tt.expectedCount)
			}
		})
	}
}

func TestInteractorImpl_GetMyVoucher(t *testing.T) {
	th := InitTest()
	initPromotion(th.DBPromotion)

	//defer th.DBPromotion.Drop(helpers.ContextWithTimeOut())

	type args struct {
		offset  int64
		user_id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFN  func()
		err     error
	}{
		{
			args: args{
				user_id: "namle12345",
				offset:  0,
			},
			mockFN: func() {
				th.MockUser.On("FindMerchantByCode", mock.AnythingOfType("string")).Return(&entities.Merchant{
					Id: "test",
				}, nil).Once()
			},
			name:    "case get my voucher",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFN()
			got, err := th.Interactor.GetMyVoucher(tt.args.offset, tt.args.user_id)
			if err != nil && tt.wantErr == false {
				t.Errorf("error %s", err.Error())
			}
			t.Log("got ... ", got)
		})
	}

}

func TestInteractorImpl_GetList(t *testing.T) {
	th := InitTest()
	initPromotion(th.DBPromotion)

	defer th.DBPromotion.Drop(helpers.ContextWithTimeOut())
	type args struct {
		offset  int64
		limit   int64
		user_id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFN  func()
		err     error
	}{
		{
			args: args{
				offset:  0,
				limit:   10,
				user_id: "namle12345",
			},
			mockFN: func() {
				th.MockUser.On("FindMerchantByCode", mock.AnythingOfType("string")).Return(&entities.Merchant{
					Id: "test",
				}, nil)
				th.MockUser.On("FindById", mock.AnythingOfType("string")).Return(&entities.User{
					Id:     "test1",
					Source: "app",
				}, nil).Once()
			},
			name:    "case get LIST voucher App User",
			wantErr: false,
		},
		{
			args: args{
				offset:  0,
				limit:   10,
				user_id: "20201015-US5F8847646CDF3461149",
			},
			mockFN: func() {
				th.MockUser.On("FindMerchantByCode", mock.AnythingOfType("string")).Return(&entities.Merchant{
					Id: "test",
				}, nil)
				th.MockUser.On("FindById", mock.AnythingOfType("string")).Return(&entities.User{
					Id:     "20201015-US5F8847646CDF3461149",
					Source: "GAPO",
				}, nil).Once()
			},
			name:    "case get LIST voucher GAPO user",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFN()
			got, err := th.Interactor.GetListVoucher(&service_promotion.GetListPromotionRequest{
				Type:   "",
				UserId: tt.args.user_id,
				Offset: tt.args.offset,
				IdGapo: "",
			})
			if err != nil && tt.wantErr == false {
				t.Errorf("error %s", err.Error())
			}
			t.Log("got LIST... ", got[0].SourceUser)
		})
	}

}

func TestInteractorImpl_UseVoucher(t *testing.T) {
	th := InitTest()
	initPromotionWithRealSOF(th.DBPromotion)

	defer th.DBPromotion.Drop(helpers.ContextWithTimeOut())

	type args struct {
		voucher_id          string
		code                string
		user_id             string
		total               int64
		amount              int64
		service_code        string
		current_balance     int64
		expected_count_used int64
		sourceOfFund        string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name:    "Voucher",
			wantErr: true,
			args: args{
				code:            "VOUCHER_11",
				user_id:         "namle11",
				total:           1,
				amount:          100,
				service_code:    "",
				current_balance: 2000000,
				sourceOfFund:    constant.SOURCE_OF_FUND_BALANCE_WALLET,
			},
			err: fmt.Errorf("Giá trị giao dịch tối thiểu để áp dụng mã khuyến mãi là %v", 10000),
		},
		{
			name:    "ko đủ tiền thực hiện giao dịch",
			wantErr: true,
			args: args{
				code:            "VOUCHER_11",
				user_id:         "namle11",
				total:           1,
				amount:          10000,
				service_code:    "",
				current_balance: 100,
				sourceOfFund:    constant.SOURCE_OF_FUND_BALANCE_WALLET,
			},
			err: fmt.Errorf(errors.NotFoundWallet().Error()),
		},
		{
			name:    "Check count use voucher per all",
			wantErr: false,
			args: args{
				code:                "VOUCHER_11",
				user_id:             "namle11",
				total:               1,
				amount:              10000,
				service_code:        "",
				current_balance:     100000,
				expected_count_used: 1,
				voucher_id:          "voucher11",
				sourceOfFund:        constant.SOURCE_OF_FUND_BALANCE_WALLET,
			},
		},
		{
			name:    "Voucher with sourceOfFund",
			wantErr: true,
			args: args{
				code:            "VOUCHER_12",
				user_id:         "namle11",
				total:           1,
				amount:          10000,
				service_code:    "",
				current_balance: 200,
				voucher_id:      "voucher12",
				sourceOfFund:    constant.SOURCE_OF_FUND_BALANCE_WALLET,
			},
			err: fmt.Errorf(errors.NotFoundWallet().Error()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &service_promotion.UseVoucherResponse{
				LastAmount:     0,
				DiscountAmount: 1000,
			}

			_, err := th.Interactor.UseVoucher(&service_promotion.UseVoucherRequest{
				Code:           tt.args.code,
				UserId:         tt.args.user_id,
				TraceId:        "trace_id....",
				Total:          tt.args.total,
				Amount:         tt.args.amount,
				ServiceCode:    "",
				CurrentBalance: tt.args.current_balance,
				SourceOfFund:   tt.args.sourceOfFund,
			}, res)

			if err != nil && tt.wantErr == false {
				t.Errorf("error %s", err.Error())
			}
			assert.Equal(t, err, tt.err)
			if err == nil {
				filter := bson.M{"user_id": tt.args.user_id, "voucher_id": tt.args.voucher_id, "action": "u"}
				gotCountBuyVoucher, _ := th.DBPromotion.Collection("promo_log_cmd").CountDocuments(helpers.ContextWithTimeOut(), filter)
				assert.Equal(t, gotCountBuyVoucher, tt.args.expected_count_used)
			}

		})
	}
}
