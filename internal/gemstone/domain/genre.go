package domain

type Genres struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;index;"`
	GenreType string `json:"name,omitempty" gorm:"foreignKey:GenreType;references:GenreType;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
