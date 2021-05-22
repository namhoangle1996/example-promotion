package main

import (
	"github.com/joho/godotenv"
	"gitlab.com/wallet-gpay/service-promotion/configs"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/interactors"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/presenters/grpc"
)

func init() {
	godotenv.Load()
}


func main() {
	// New Service

	config, err := configs.LoadConfig()

	if err != nil {
		panic(err)
	}

	grpc.NewHandler(
		interactors.Init(config),
	).StartServer()
}
