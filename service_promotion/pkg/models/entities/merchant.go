package entities

type Merchant struct {
	Id          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Email       string `json:"email,omitempty" bson:"email,omitempty"`
	Code        string `json:"code,omitempty" bson:"code,omitempty"`
	Logo        string `json:"logo,omitempty" bson:"logo,omitempty"`
	SecretKey   string `json:"secret_key,omitempty" bson:"secret_key,omitempty"`
	Address     string `json:"address,omitempty" bson:"address,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Website     string `json:"website,omitempty" bson:"website,omitempty"`
}
