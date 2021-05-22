package entities

type PopUp struct {
	Id          string `json:"_id" bson:"_id"`
	Type        string `json:"type" bson:"type"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	ImageURL    string `json:"image_url" bson:"image_url"`
	StartDate   int    `json:"start_date" bson:"start_date"`
	EndDate     int    `json:"end_date" bson:"end_date"`
	Frequency   int    `json:"frequency" bson:"frequency"`
	ExternalURL string `json:"external_url" bson:"external_url"`
	AllowAll    bool   `json:"allow_all" bson:"allow_all"`
}
