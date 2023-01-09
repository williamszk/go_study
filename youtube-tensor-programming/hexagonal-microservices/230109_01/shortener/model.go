package shortener

type Redirect struct {
	Code      string `json:"code" bson:"code" msgpack:"code"`
	URL       string  `json:"url" bson:"url" msgpack:"url" validate:"empty=false & format=url"`
	CreatedAt int64  `json:"created_at" bson:"create_at" msgpack:"create_at"`
}
