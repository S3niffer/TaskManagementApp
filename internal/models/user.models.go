package models

type User struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password_hash string `json:"password_hash"`
	Create_at     string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
}

type NewUser struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password_hash string `json:"-"`
	Create_at     string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	AccessToken   string `json:"access_token"`
}
