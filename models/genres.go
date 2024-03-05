package models

const GenreTableName = "genres"

func (*GenreBase) TableName() string {
	return GenreTableName
}

type GenreBase struct {
	BaseModel          `json:",omitempty"`
	CreateGenreRequest `json:",omitempty"`
	GenID              int `gorm:"column:gen_id;primaryKey" json:"gen_id,omitempty"`
}

type CreateGenreRequest struct {
	GenTitle string `gorm:"column:gen_title" json:"gen_title,omitempty"`
}

type Genre struct {
	GenreBase  GenreBase    `gorm:"embedded" json:",omitempty"`
	MovieGenre []MovieGenre `gorm:"foreignKey:GenID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
}
