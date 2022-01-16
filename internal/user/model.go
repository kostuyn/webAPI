package user

type User struct {
	Id           string `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email" bson:"email"`
	UserName     string `json:"userName" bson:"userName"`
	PasswordHash string `json:"-" bson:"passwordHash"`
}

type CreateUserDto struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
