package model

type Usergroup struct {
	//Id      int64 `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	//UserId  int64 `json:"id" column:"user_id"`
	
	Id int64 `json:"id" column:"group_id"`
	UserId  int64 `json:"user_id" column:"user_id"`
}

func (usergroup *Usergroup) Table() string {
	return "userGroup"
}

func (usergroup *Usergroup) String() string {
	return Stringify(usergroup)
}
