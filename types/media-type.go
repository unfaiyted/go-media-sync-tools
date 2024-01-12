package types

import "github.com/google/uuid"

type ConfiguredLibrary struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID   uuid.UUID
	ConfigID uuid.UUID
}

type SyncType struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;"`
}

type Config struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	User      *User     `gorm:"foreignKey:ID"`
	UserId    string
	Clients   []*ConfiguredClient  `gorm:"foreignKey:ConfigID"`
	Libraries []*ConfiguredLibrary `gorm:"foreignKey:ConfigID"`
	Sync      *SyncType            `gorm:"foreignKey:ID"`
}
