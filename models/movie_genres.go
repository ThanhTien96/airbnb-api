package models

type MovieGenre struct {
	MovID int    `gorm:"column:mov_id" json:"mov_id,omitempty"`
	Movie *Movie `gorm:"foreignKey:MovID" json:",omitempty"`
	GenID int    `gorm:"column:gen_id" json:"gen_id,omitempty"`
	Genre *Genre `gorm:"foreignKey:GenID" json:",omitempty"`
}
