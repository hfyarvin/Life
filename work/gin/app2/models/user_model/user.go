package user_model

import ()

type User struct {
	User     string `from:"user" json:"user"`
	Password string `from:"pwd" json:"pwd"`
}
