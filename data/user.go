package data

type User struct {
	ID             string `json:"id" bson:"_id"`
	Username       string `json:"username" bson:"username"`
	DisplayName    string `json:"display_name" bson:"display_name"`
	Email          string `json:"email" bson:"email"`
	ProfilePicture string `json:"pfp" bson:"pfp"`
	PasswordHash   string `json:"pfp" bson:"password_hash"`
	PasswordSalt   string `json:"pfp" bson:"password_salt"`
	Timestamp      int64  `json:"timestamp" bson:"timestamp"`
}
