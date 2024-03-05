package models

type Director struct {
	DirID          int              `gorm:"column:dir_id;primaryKey" json:"dir_id,omitempty"`
	DirFName       string           `gorm:"column:dir_fname" json:"dir_fname,omitempty"`
	DirLName       string           `gorm:"column:dir_lname" json:"dir_lname,omitempty"`
	MovieDirection []MovieDirection `gorm:"foreignKey:DirID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
}
