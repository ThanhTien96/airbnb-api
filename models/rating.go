package models

type Rating struct {
	MovID       int       `gorm:"column:mov_id" json:"mov_id,omitempty"`
	Movie       *Movie    `gorm:"foreignKey:MovID" json:",omitempty"`
	RevID       int       `gorm:"column:rev_id" json:"rev_id,omitempty"`
	Reviewer    *Reviewer `gorm:"foreignKey:RevID" json:",omitempty"`
	RevStars    int       `gorm:"column:rev_stars" json:"rev_stars,omitempty"`
	NumORatings int       `gorm:"column:num_o_ratings" json:"num_o_ratings,omitempty"`
}
