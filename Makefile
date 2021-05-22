BINARY=engine
IMAGE=registry.gitlab.com/wallet-gpay/service-promotion
GOPATH:=$(shell go env GOPATH)

docker-dev:
	rm -rf vendor
	sh ./build.sh buildNumber dev- ${IMAGE}

engine:
	go build -o ${BINARY} ./cmd/api/main.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

gen-proto-macos:
	# gen for service_transaction
	rm -rf ./proto/service_transaction && protoc --proto_path=./protobufs --micro_out=./proto --go_out=./proto ./protobufs/service_transaction/*.proto

	# gen for service_statistic
	rm -rf ./proto/service_statistic && protoc --proto_path=./protobufs --micro_out=./proto --go_out=./proto ./protobufs/service_statistic/*.proto

	# gen for service_merchant_fee
	rm -rf ./proto/service_merchant_fee && protoc --proto_path=./protobufs --micro_out=./proto --go_out=./proto ./protobufs/service_merchant_fee/*.proto

	rm -rf ./proto/service_promotion && protoc --proto_path=./protobufs --micro_out=./proto --go_out=./proto ./protobufs/service_promotion/*.proto

gen-proto:
	# gen for service_transaction
	rm -rf ./proto/service_transaction && protoc --proto_path=./protobufs --micro_out=./proto --go_out=./proto ./protobufs/service_transaction/*.proto && sed -i -e "s/,omitempty//g" ./proto/service_transaction/*.pb.go

	# gen for service_statistic
	rm -rf ./proto/service_statistic && protoc --proto_path=./protobufs --micro_out=./proto --go_out=./proto ./protobufs/service_statistic/*.proto && sed -i -e "s/,omitempty//g" ./proto/service_statistic/*.pb.go

	# gen for service_merchant_fee
	rm -rf ./proto/service_merchant_fee && protoc --proto_path=./protobufs --micro_out=./proto --go_out=./proto ./protobufs/service_merchant_fee/*.proto && sed -i -e "s/,omitempty//g" ./proto/service_merchant_fee/*.pb.go

	rm -rf ./proto/service_promotion && protoc --proto_path=./protobufs --micro_out=./proto --go_out=./proto ./protobufs/service_promotion/*.proto && sed -i -e "s/,omitempty//g" ./proto/service_promotion/*.pb.go

gen-mocks:
	echo "generate mocks"
	cd ./service_promotion/pkg/repositories/data_service && mockery --name=IServiceUser
	cd ./service_promotion/pkg/repositories/datasource/mongodb && mockery --name=IVoucher && mockery --name=IPopup && mockery --name=IUserHasVoucher && mockery --name=IUserApplyVoucher  && mockery --name=ILogVoucher

test:
	go test -mod vendor  -p 1 `go list ./... | grep -v /vendor/`