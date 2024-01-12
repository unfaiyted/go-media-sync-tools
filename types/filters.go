package types

type FilterType int

const (
	FilterTypeUnknown FilterType = iota
	FilterTypeEmby
	FilterTypeJellyfin
	FilterTypePlex
	FilterTypeTrakt
	FilterTypeTmdb
	FilterTypeMDB
)

type Filter struct {
	ID          string
	typeID      string
	mediaListID string
	label       string
	value       string
}

type FilterTypes struct {
	ID       string
	ClientId string
	Name     string
	Type     FilterType
	Label    string
}

type Filters interface {
	GetType() FilterType
}

// func ProcessFilter(f Filters) {
//   switch f.GetType() {
//   case FilterTypeEmby:
// }

type BaseFilters struct {
	ID       string
	ClientID string
	Type     FilterType
}

type EmbyFilters struct {
	BaseFilters
	ListID                *string
	Search                *string
	Limit                 *int
	Offset                *int
	SortBy                *string
	SortOrder             *string
	GroupProgramsBySeries *bool
	Recursive             *bool
	IncludeItemTypes      *string
	StudioIDs             *string
	StudioNames           *string
	Is3D                  *bool
	OfficialRatings       *string
	HasTrailer            *bool
	HasSpecialFeature     *bool
	HasThemeSong          *bool
	HasThemeVideo         *bool
	HasOverview           *bool
	HasImdbID             *bool
	HasTmdbID             *bool
	HasTvdbID             *bool
	AudioLanguages        *string
	VideoCodecs           *string
	AudioCodecs           *string
	IsPlayed              *bool
	IsFavorite            *bool
}

type PlexFilters struct {
	BaseFilters
	ListID        *string
	Library       *string
	Type          *string
	Year          *int
	Decade        *int
	Genre         *string
	Country       *string
	Studio        *string
	Actor         *string
	Director      *string
	ContentRating *string
	Resolution    *string
	AudioChannels *string
	AudioCodec    *string
	VideoCodec    *string
	Sort          *string
	Title         *string
	Limit         *int
	Offset        *int
}

type TraktFilters struct {
	BaseFilters
	ListID   *string
	ListSlug *string
}

type TmdbFilters struct {
	BaseFilters
	ListID                *string
	Page                  *int
	Language              *string
	PrimaryReleaseYear    *int
	Region                *string
	ReleaseDateGte        *string
	ReleaseDateLte        *string
	SortBy                *string
	VoteCountGte          *int
	VoteCountLte          *int
	VoteAverageGte        *int
	VoteAverageLte        *int
	WithCast              *string
	WithCrew              *string
	WithKeywords          *string
	WithRuntimeGte        *int
	WithRuntimeLte        *int
	WithOriginalLanguage  *string
	WithPeople            *string
	WithCompanies         *string
	WithGenres            *string
	WithoutGenres         *string
	WithoutKeywords       *string
	WithoutCompanies      *string
	WithWatchProviders    *string
	WithoutWatchProviders *string
	Year                  *int
}

type TmdbShowFilters struct {
	TmdbFilters
	AirDateGte           *string
	AirDateLte           *string
	FirstAirDateYear     *int
	FirstAirDateGte      *string
	FirstAirDateLte      *string
	WithNetworks         *string
	IncludeAdult         *bool
	ScreenedTheatrically *bool
	Timezone             *string
	WithStatus           *string
	WithType             *string
}

type MdbFilter struct {
	BaseFilters
	ListID  *string
	Library *string
}
