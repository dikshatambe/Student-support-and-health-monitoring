package model

type UserQuery struct {
	QId       int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	Id        int64  `json:"user_id" column:"user_id"`
	Query string `json:"query" column:"query"`
}

func (userquery *UserQuery) Table() string {
	return "userQuery"
}

func (userquery *UserQuery) String() string {
	return Stringify(userquery)
}
