package types

import (
	"time"

	"github.com/google/uuid"
)

type SyncOptions struct {
	ID          string
	ConfigID    uuid.UUID
	Collections bool
	Playlists   bool
	LovedTracks bool
	TopLists    bool
	Trakt       bool
	Watched     bool
	Ratings     bool
	Libraries   bool
}

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name       string
	Username   string `gorm:"unique"`
	Email      string `gorm:"unique"`
	Password   string
	ConfigID   string
	Lists      []*MediaList `gorm:"foreignKey:CreatorID"`
	UserConfig *Config      `gorm:"foreignKey:ID"`
	CreatedAt  time.Time
}
