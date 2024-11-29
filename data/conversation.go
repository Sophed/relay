package data

type Conversation struct {
	ID        string   `json:"id" bson:"_id"`
	Members   []string `json:"members" bson:"members"`
	Direct    bool     `json:"direct" bson:"direct"`
	Timestamp int64    `json:"timestamp" bson:"timestamp"`
}
