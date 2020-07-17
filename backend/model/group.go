package model

type Group struct {
	Id        int64  `json:"id,omitempty" key:"primary" column:"id"`
	GroupName string `json:"group_name" column:"group_name"`
	GroupInfo string `json:"group_info" column:"group_info"`
}

func (group *Group) Table() string {
	return "groups1"
}

func (group *Group) String() string {
	return Stringify(group)
}
