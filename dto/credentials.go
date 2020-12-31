package dto

type Credentials struct {
	Username string `form:"username"`
	Password string `forn:"password"`
}

type JWT struct {
	Token string `json:"token"`
}
