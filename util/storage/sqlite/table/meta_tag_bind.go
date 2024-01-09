package table

type MetaTagBind struct {
	ID         uint `gorm:"primaryKey"`
	MetaId     uint `gorm:"index"`
	TagId      uint `gorm:"index"`
	CreateTime int64
}
