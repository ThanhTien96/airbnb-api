package models

type MovieCast struct {
	MovID int    `gorm:"column:mov_id" json:"mov_id,omitempty"`
	Movie *Movie `gorm:"foreignKey:MovID" json:",omitempty"`
	ActID int    `gorm:"column:act_id" json:"act_id,omitempty"`
	Actor *Actor `gorm:"foreignKey:ActID" json:",omitempty"`
	Role  string `gorm:"column:role" json:"role,omitempty"`
}
