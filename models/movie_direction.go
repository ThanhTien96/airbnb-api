package models

type MovieDirection struct {
	MovID    int       `gorm:"column:mov_id" json:"mov_id,omitempty"`
	Movie    *Movie    `gorm:"foreignKey:MovID" json:",omitempty"`
	DirID    int       `gorm:"column:dir_id" json:"dir_id,omitempty"`
	Director *Director `gorm:"foreignKey:DirID" json:",omitempty"`
}
