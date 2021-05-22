package entities

import "time"

type User struct {
	Id                           string    `json:"id" bson:"_id,omitempty"`
	IdGutina                     int64     `json:"id_gutina" bson:"id_gutina,omitempty"`
	PhoneNumber                  string    `json:"phone_number" bson:"phone_number,omitempty"`
	Email                        string    `json:"email" bson:"email,omitempty"`
	Name                         string    `json:"name"  bson:"name,omitempty"`
	Avatar                       string    `json:"avatar" bson:"avatar,omitempty"`
	Gender                       string    `json:"gender" bson:"gender,omitempty"`
	Status                       string    `json:"status" bson:"status,omitempty"`
	Timezone                     string    `json:"timezone"`
	Language                     string    `json:"language"`
	Title                        string    `json:"title"`
	DateTime                     int64     `json:"date_time" bson:"date_time,omitempty"`
	Currency                     string    `json:"currency" bson:"currency"`
	Role                         string    `json:"role" bson:"role"`
	IsFacebookConnect            bool      `json:"is_facebook_connect" bson:"is_facebook_connect"`
	EmailVerify                  bool      `json:"email_verify" bson:"email_verify"`
	CanUpdateEmail               bool      `json:"can_update_email" bson:"can_update_email"`
	CanAddBankAccount            bool      `json:"can_add_bank_account" bson:"can_add_bank_account"`
	CanUpdatePhoneNumber         bool      `json:"can_update_phone_number" bson:"can_update_phone_number"`
	PhoneNumberVerify            bool      `json:"phone_number_verify" bson:"phone_number_verify"`
	AccessToken                  string    `json:"access_token" bson:"-"`
	IsExists                     bool      `json:"is_exists" bson:"is_exists"`
	IdentityNumber               string    `json:"identity_number" bson:"identity_number"`
	BirthDay                     string    `json:"birth_day" bson:"birth_day"`
	ConnectedFacebook            bool      `json:"connected_facebook" bson:"connected_facebook"`
	Address                      string    `json:"address" bson:"address"`
	MessageTopics                []string  `json:"message_topics" bson:"message_topics"`
	DeviceIdLogin                string    `json:"device_id_login" bson:"device_id_login"`
	LastDeviceIdLogin            string    `json:"last_device_id_login" bson:"last_device_id_login"`
	Password                     string    `json:"password" bson:"password"`
	Long                         float64   `json:"long"`
	Lat                          float64   `json:"lat"`
	OriginDevice                 string    `json:"origin_device" bson:"origin_device"`
	Source                       string    `json:"source"`
	RandomString                 string    `json:"random_string" bson:"random_string"`
	LastAmountNetPrimaryWallet   int64     `json:"last_amount_net_primary_wallet" bson:"last_amount_net_primary_wallet"`
	LastAmountPrimaryWallet      int64     `json:"last_amount_primary_wallet" bson:"last_amount_primary_wallet"`
	IncrementAmountPrimaryWallet int64     `json:"increment_amount_primary_wallet" bson:"increment_amount_primary_wallet"`
	DecrementAmountPrimaryWallet int64     `json:"decrement_amount_primary_wallet" bson:"decrement_amount_primary_wallet"`
	LastReadAllMessage           time.Time `json:"last_read_all_message" bson:"last_read_all_message"`
	CreatedAt                    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt                    time.Time `json:"deleted_at" bson:"deleted_at"`
}
