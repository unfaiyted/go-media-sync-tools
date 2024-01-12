package types

import "github.com/google/uuid"

type LibraryType int

const (
	LibraryTypeUnknown LibraryType = iota
	LibraryTypeMovies
	LibraryTypeShows
	LibraryTypeMusic
	LibraryTypeGames
	LibraryTypeBooks
	LibraryTypeAudiobooks
	LibraryTypeAnime
	// ... other types ...
)

// This represents a library as a type of library
type Library struct {
	ID       uuid.UUID
	Name     string
	Type     LibraryType
	Clients  []*LibraryClient `gorm:"foreignKey:LibraryID"`
	ConfigID uuid.UUID
	Lists    *MediaList `gorm:"foreignKey:LibraryID"`
}

// This represents an instance of a library that has beenc configured by a user
type LibraryClient struct {
	ID          uuid.UUID
	LibraryID   uuid.UUID
	ClientID    uuid.UUID
	MediaListID *uuid.UUID
	LibraryName string
}
