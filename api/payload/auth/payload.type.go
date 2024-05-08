package payload

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Avatar string `json:"avatar"`
}
