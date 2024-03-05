package models

type Reviewer struct {
	RevID   int      `gorm:"column:rev_id;primaryKey" json:"rev_id,omitempty"`
	RevName string   `gorm:"column:rev_name" json:"rev_name,omitempty"`
	Rating  []Rating `gorm:"foreignKey:RevID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
}
