package model

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Profile struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Bio   string `json:"bio"`
}

type UserObject struct {
	UserObject string `json:"userObject"`
}
