package model

type User struct {
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserPass  string `json:"user_pass"`
	UserType  string `json:"user_type"`
}

type Credentials struct {
	UserName string `json:"user_name"`
	UserPass string `json:"user_pass"`
}
