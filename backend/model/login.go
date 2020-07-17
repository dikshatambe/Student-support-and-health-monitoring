package model

type Login struct {
	Userid   int64  `json:"id" column:"id"`
	Password string `json:"password" column:"password"`
}

func (login *Login) Table() string {
	return "users"
}

func (login *Login) String() string {
	return Stringify(login)
}
