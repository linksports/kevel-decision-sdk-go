package model

type User struct {
	Key string `json:"key"`
}

func NewUser(key string) User {
	return User{key}
}
