package models

type Actor struct {
	ActID     int         `gorm:"column:act_id;primaryKey" json:"act_id,omitempty"`
	ActFName  string      `gorm:"column:act_fname" json:"act_fname,omitempty"`
	ActLName  string      `gorm:"column:act_lname" json:"act_lname,omitempty"`
	ActGender string      `gorm:"column:act_gender" json:"act_gender,omitempty"`
	MovieCast []MovieCast `gorm:"foreignKey:ActID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
}
