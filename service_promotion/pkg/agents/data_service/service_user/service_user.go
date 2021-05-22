package service_user

import (
	"github.com/micro/go-micro/v2"
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"gitlab.com/wallet-gpay/service-promotion/proto/service_user"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/aggregates"
	"gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

func convertUserDetailDTO2TypeUser(dto service_user.UserDetailDTO) *aggregates.UserDetail {

	balances := make(map[int]entities.Balances, len(dto.Balances))

	for k, v := range dto.Balances {
		balances[k] = entities.Balances{
			Id:              v.Id,
			AmountAvailable: v.AmountAvailable,
			AmountFreeze:    v.AmountFreeze,
			Type:            v.Type,
			UserId:          v.UserId,
			FreezeIds:       v.FreezeIds,
			Currency:        v.Currency,
		}
	}

	return &aggregates.UserDetail{
		User: entities.User{
			Id:                           dto.User.Id,
			IdGutina:                     dto.User.IdGutina,
			PhoneNumber:                  dto.User.PhoneNumber,
			Email:                        dto.User.Email,
			Name:                         dto.User.Name,
			Avatar:                       dto.User.Avatar,
			Gender:                       dto.User.Gender,
			Status:                       dto.User.Status,
			Timezone:                     dto.User.Timezone,
			Language:                     dto.User.Language,
			Title:                        dto.User.Title,
			DateTime:                     dto.User.DateTime,
			Currency:                     dto.User.Currency,
			Role:                         dto.User.Role,
			IsFacebookConnect:            dto.User.IsFacebookConnect,
			EmailVerify:                  dto.User.EmailVerify,
			CanUpdateEmail:               dto.User.CanUpdateEmail,
			CanAddBankAccount:            dto.User.CanAddBankAccount,
			CanUpdatePhoneNumber:         dto.User.CanUpdatePhoneNumber,
			PhoneNumberVerify:            dto.User.PhoneNumberVerify,
			AccessToken:                  dto.User.AccessToken,
			IsExists:                     dto.User.IsExists,
			IdentityNumber:               dto.User.IdentityNumber,
			BirthDay:                     dto.User.BirthDay,
			ConnectedFacebook:            dto.User.ConnectedFacebook,
			Address:                      dto.User.Address,
			MessageTopics:                dto.User.MessageTopics,
			DeviceIdLogin:                dto.User.DeviceIdLogin,
			LastDeviceIdLogin:            dto.User.LastDeviceIdLogin,
			OriginDevice:                 dto.User.OriginDevice,
			Source:                       dto.User.Source,
			RandomString:                 dto.User.RandomString,
			LastAmountNetPrimaryWallet:   dto.User.LastAmountNetPrimaryWallet,
			LastAmountPrimaryWallet:      dto.User.LastAmountPrimaryWallet,
			IncrementAmountPrimaryWallet: dto.User.IncrementAmountPrimaryWallet,
			DecrementAmountPrimaryWallet: dto.User.DecrementAmountPrimaryWallet,
		},
		Balances: balances,
	}
}

func convertUserDTO2TypeUser(dto service_user.UserDTO) *entities.User {
	return &entities.User{
		Id:                           dto.Id,
		IdGutina:                     dto.IdGutina,
		PhoneNumber:                  dto.PhoneNumber,
		Email:                        dto.Email,
		Name:                         dto.Name,
		Avatar:                       dto.Avatar,
		Gender:                       dto.Gender,
		Status:                       dto.Status,
		Timezone:                     dto.Timezone,
		Language:                     dto.Language,
		Title:                        dto.Title,
		DateTime:                     dto.DateTime,
		Currency:                     dto.Currency,
		Role:                         dto.Role,
		IsFacebookConnect:            dto.IsFacebookConnect,
		EmailVerify:                  dto.EmailVerify,
		CanUpdateEmail:               dto.CanUpdateEmail,
		CanAddBankAccount:            dto.CanAddBankAccount,
		CanUpdatePhoneNumber:         dto.CanUpdatePhoneNumber,
		PhoneNumberVerify:            dto.PhoneNumberVerify,
		AccessToken:                  dto.AccessToken,
		IsExists:                     dto.IsExists,
		IdentityNumber:               dto.IdentityNumber,
		BirthDay:                     dto.BirthDay,
		ConnectedFacebook:            dto.ConnectedFacebook,
		Address:                      dto.Address,
		MessageTopics:                dto.MessageTopics,
		DeviceIdLogin:                dto.DeviceIdLogin,
		LastDeviceIdLogin:            dto.LastDeviceIdLogin,
		OriginDevice:                 dto.OriginDevice,
		Source:                       dto.Source,
		RandomString:                 dto.RandomString,
		LastAmountNetPrimaryWallet:   dto.LastAmountNetPrimaryWallet,
		LastAmountPrimaryWallet:      dto.LastAmountPrimaryWallet,
		IncrementAmountPrimaryWallet: dto.IncrementAmountPrimaryWallet,
		DecrementAmountPrimaryWallet: dto.DecrementAmountPrimaryWallet,
	}
}

type repositoryImpl struct {
	Service service_user.Service
}

func (r repositoryImpl) FindDetailById(id string) (*aggregates.UserDetail, error) {
	resp, err := r.Service.FindUserDetailById(helpers.ContextWithTimeOut(), &service_user.FindUserDetailByIdRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return convertUserDetailDTO2TypeUser(*resp.UserDetail), nil
}

func (r repositoryImpl) FindById(id string) (*entities.User, error) {
	resp, err := r.Service.FindUserById(helpers.ContextWithTimeOut(), &service_user.FindUserByIdRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return convertUserDTO2TypeUser(*resp.User), nil
}

func (r repositoryImpl) FindMerchantByCode(code string) (*entities.Merchant, error) {
	resp, err := r.Service.FindMerchantAccountByCode(helpers.ContextWithTimeOut(), &service_user.FindMerchantAccountByCodeRequest{
		Code: code,
	})

	if err != nil {
		return nil, err
	}

	return &entities.Merchant{
		Id:          resp.MerchantDetail.Merchant.Id,
		Email:       resp.MerchantDetail.Merchant.Email,
		Code:        resp.MerchantDetail.Merchant.Code,
		Logo:        resp.MerchantDetail.Merchant.Logo,
		SecretKey:   resp.MerchantDetail.Merchant.SecretKey,
		Address:     resp.MerchantDetail.Merchant.Address,
		Name:        resp.MerchantDetail.Merchant.Name,
		PhoneNumber: resp.MerchantDetail.Merchant.PhoneNumber,
		Website:     resp.MerchantDetail.Merchant.Website,
	}, nil
}

func (r repositoryImpl) FindMerchantById(id string) (*entities.Merchant, error) {
	resp, err := r.Service.FindMerchantAccountByID(helpers.ContextWithTimeOut(), &service_user.FindMerchantAccountByIdRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return &entities.Merchant{
		Id:          resp.MerchantDetail.Merchant.Id,
		Email:       resp.MerchantDetail.Merchant.Email,
		Code:        resp.MerchantDetail.Merchant.Code,
		Logo:        resp.MerchantDetail.Merchant.Logo,
		SecretKey:   resp.MerchantDetail.Merchant.SecretKey,
		Address:     resp.MerchantDetail.Merchant.Address,
		Name:        resp.MerchantDetail.Merchant.Name,
		PhoneNumber: resp.MerchantDetail.Merchant.PhoneNumber,
		Website:     resp.MerchantDetail.Merchant.Website,
	}, nil
}

func NewRepositoryImpl(serviceName string) *repositoryImpl {
	service := micro.NewService()
	service.Init()

	serviceCore := service_user.NewService(serviceName, service.Client())

	return &repositoryImpl{
		Service: serviceCore,
	}
}
