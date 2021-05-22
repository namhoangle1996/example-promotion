package interactors

import (
	"gitlab.com/wallet-gpay/service-promotion/configs"
	"gitlab.com/wallet-gpay/service-promotion/gpooling"
	"gitlab.com/wallet-gpay/service-promotion/helpers/saga"
	"gitlab.com/wallet-gpay/service-promotion/logger"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/data_service/service_user"
	infrastructureMongo "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo/log_voucher"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo/popup"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo/promo_user_has_voucher"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo/promo_users_apply_voucher"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/agents/datasource/mongo/vouchers"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/repositories"
	"os"
)

type InteractorImpl struct {
	*repositories.HandlerRepository
	gpooling.IPool
	logger.ILogger
	LogSaga saga.Store
}

func Init(config *configs.Config) *InteractorImpl {

	primaryDb := infrastructureMongo.NewMongoDBconnection(config.MongoURI)

	pool, err := gpooling.NewPooling(config.MaxPoolSize)
	if err != nil {
		panic(err.Error())
	}

	hostname, _ := os.Hostname()
	env := config.ENV
	log, err := logger.NewLogger(hostname, env)

	if err != nil {
		panic(err.Error())
	}

	return &InteractorImpl{
		HandlerRepository: repositories.NewHandlerRepository(
			vouchers.NewRepository(primaryDb),
			log_voucher.NewRepository(primaryDb),
			promo_users_apply_voucher.NewRepository(primaryDb),
			promo_user_has_voucher.NewRepository(primaryDb),
			service_user.NewRepositoryImpl(config.Services.ServiceUser),
			popup.NewRepository(primaryDb),
		),
		IPool:   pool,
		ILogger: log,
		LogSaga:               saga.New(),
	}
}
