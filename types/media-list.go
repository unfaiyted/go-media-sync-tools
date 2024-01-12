package types

import (
	"time"

	"github.com/google/uuid"
)

type MediaListType int

const (
	MediaListTypeUnknown MediaListType = iota
	MediaListTypeCollection
	MediaListTypeLibrary
)

type Provider int

const (
	ProviderUnknown Provider = iota
	ProviderOpenAI
	ProviderMDB
	ProviderTrakt
)

type MediaType int

const (
	MediaTypeUnknown MediaType = iota
	MediaTypeMovie
	MediaTypeShow
	MediaTpeEpisode
	MediaTypeSeason
)

type MediaListOptions struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID         string
	Type           MediaListType
	Clients        []*ConfiguredClient
	MediaListID    string
	SyncLibraryID  string
	sync           bool
	UpdateImages   bool
	deleteExisting bool
}

type MediaList struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;"`
	sourceID    string
	ClientID    string
	PosterID    string
	CreatorID   string
	Name        string
	SortName    string
	Description *string
	CreatedAt   time.Time
	Type        MediaListType
	// Filters     Filters          `gorm:"foreignKey:FiltersID"`
	// Items   []*MediaListItem `gorm:"foreignKey:MediaListID;references:ID"`
	Creator User           `gorm:"foreignKey:ID"`
	Poster  ProviderPoster `gorm:"foreignKey:ID"`
}

type MediaProviderIDs struct {
	imdb           *string
	tvdb           *string
	tmdb           *string
	trakt          *string
	tvRage         *string
	rottenTomatoes *string
	aniList        *string
}

type MediaItemRatings struct {
	tmdb           *float32
	imdb           *float32
	trakt          *float32
	metacritic     *float32
	rottenTomatoes *float32
	tvdb           *float32
	aniList        *float32
}

type MediaItem struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Title          string
	Year           string
	Type           MediaType
	SortTitle      string
	OriginalTitle  string
	Tagline        string
	Poster         string
	Description    string
	ParentalRating string
	Genres         []*string
	ReleaseDate    time.Time
	DateAdded      time.Time
	Providers      MediaProviderIDs
	Ratings        MediaItemRatings
}

type MediaListItem struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;"`
	MediaListID   string
	MediaItemID   string
	Poster        MediaPoster
	MediaPosterID string
	Item          MediaItem
	SourceId      string
	DateAdded     time.Time
}
