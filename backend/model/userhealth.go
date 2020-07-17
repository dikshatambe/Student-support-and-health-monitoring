package model

type UserHealth struct {
	Id        int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	UId       int64  `json:"user_id" column:"user_id"`
	Age       int64  `json:"age" column:"age"`
	Temp      int64  `json:"temp" column:"temp"`
	SourThroat bool `json:"sour_throat" column:"sour_throat"`
	MedicalHistory string `json:"medical_history" column:"medical_history"`
}

func (userhealth *UserHealth) Table() string {
	return "userHealth"
}

func (userhealth *UserHealth) String() string {
	return Stringify(userhealth)
}
