package domain

type Link struct {
	ID        int    `gorm:"primaryKey"`
	URL       string `gorm:"unique"`
	ShortURL  string `gorm:"unique"`
	Increment int    `gorm:"default:0"`
	QrCode    string `gorm:"default:''"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
}
