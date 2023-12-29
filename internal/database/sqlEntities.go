package database

import (
	"time"

	"github.com/google/uuid"
)

type DeviceEntity struct {
	DeviceUuid uuid.UUID `gorm:"type:uuid;primaryKey;column:device_uuid"`

	Longitude float64 `gorm:"type:decimal(10,8)"`
	Latitude  float64 `gorm:"type:decimal(10,8)"`
}

func (table DeviceEntity) TableName() string {
	return "devices"
}

type AdEntity struct {
	AdUuid uuid.UUID `gorm:"type:uuid;primaryKey;column:ad_uuid"`

	Name     string    `gorm:"type:varchar;column:name"`
	CreateAt time.Time `gorm:"column:created_at"`

	Devices []DeviceEntity `gorm:"many2many:ads_devices;"`
}

func (table AdEntity) TableName() string {
	return "ads"
}

type UserEntity struct {
	UserUuid uuid.UUID `gorm:"type:uuid;primaryKey;column:user_uuid"`

	Name string `gorm:"type:varchar;column:name"`

	Ads []AdEntity `gorm:"many2many:users_ads;"`
}

func (table UserEntity) TableName() string {
	return "users"
}
