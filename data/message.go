package data

type Message struct {
	ID        string `json:"id" bson:"_id"`
	Content   string `json:"content" bson:"content"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}
