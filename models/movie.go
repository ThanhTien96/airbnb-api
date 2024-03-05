package models

const MovieTableName = "movies"

func (*BaseModel) TableName() string {
	return MovieTableName
}

type MovieBase struct {
	BaseModel          `json:",omitempty"`
	CreateMovieRequest `json:",omitempty"`
	MovID              int `gorm:"column:mov_id;primaryKey" json:"mov_id,omitempty"`
}

type CreateMovieRequest struct {
	MovTitle      string `gorm:"column:mov_title;not null" json:"mov_title,omitempty"`
	MovYear       int    `gorm:"column:mov_year" json:"mov_year,omitempty"`
	MovTime       int    `gorm:"column:mov_time" json:"mov_time,omitempty"`
	MovLang       string `gorm:"column:mov_lang" json:"mov_lang,omitempty"`
	MovDtRel      int64  `gorm:"column:mov_dt_rel" json:"mov_dt_rel,omitempty"`
	MovRelCountry string `gorm:"column:mov_rel_country" json:"mov_rel_country,omitempty"`
}

type Movie struct {
	MovieBase      MovieBase        `gorm:"embedded" json:",omitempty"`
	Rating         []Rating         `gorm:"foreignKey:MovID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
	MovieGenre     []MovieGenre     `gorm:"foreignKey:MovID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
	MovieDirection []MovieDirection `gorm:"foreignKey:MovID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
	MovieCast      []MovieCast      `gorm:"foreignKey:MovID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
}
