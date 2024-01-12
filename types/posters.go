package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type MediaPoster struct {
	ID         string `gorm:"primaryKey"`
	ItemId     string
	Url        string
	Type       *MediaImageType
	Overlays   []*MediaPosterOverlayOptions `gorm:"type:json"`
	Size       Size                         `gorm:"type:json"`
	Border     *MediaPosterBorderOptions    `gorm:"type:json"`
	Text       *MediaPosterTextOptions      `gorm:"type:json"`
	Gradient   *MediaPosterGradientOptions  `gorm:"type:json"`
	Background MediaPosterBackgroundOptions `gorm:"type:json"`
	Icon       *MediaPosterIconOptions      `gorm:"type:json"`
}

type MediaPosterBorderOptions struct {
	Enabled bool
	Size    Size      `gorm:"type:json"`
	Color   RGBAColor `gorm:"type:json"`
}

func (o MediaPosterBorderOptions) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *MediaPosterBorderOptions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to type assert MediaPosterBorderOptions")
	}
	return json.Unmarshal(bytes, &o)
}

type MediaPosterGradientOptions struct {
	Enabled bool
	Opacity float32
	Type    string
	Angle   int
	Colors  []*RGBAColor `gorm:"type:json"`
}

func (o MediaPosterGradientOptions) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *MediaPosterGradientOptions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to type assert MediaPosterGradientOptions")
	}
	return json.Unmarshal(bytes, &o)
}

type MediaPosterShadowOptions struct {
	Enabled      bool
	Color        *RGBAColor `gorm:"type:json"`
	Offset       *int
	Blur         *int
	Transparency *int
}

func (o MediaPosterShadowOptions) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *MediaPosterShadowOptions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to type assert MediaPosterShadowOptions")
	}
	return json.Unmarshal(bytes, &o)
}

type MediaPosterTextOptions struct {
	Enabled  bool
	Text     *string
	Font     *string
	Position *Coord     `gorm:"type:json"`
	Color    *RGBAColor `gorm:"type:json"`
	Border   *MediaPosterBorderOptions
	Shadow   *MediaPosterShadowOptions
}

func (o MediaPosterTextOptions) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *MediaPosterTextOptions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to type assert MediaPosterTextOptions")
	}
	return json.Unmarshal(bytes, &o)
}

type MediaPosterBackgroundOptions struct {
	Enabled bool
	Url     *string
	// # image Optional[Image] = Nong
	Color    *RGBAColor `gorm:"type:json"`
	Position *Coord     `gorm:"type:json"`
	Opacity  *float32
	Border   *MediaPosterBorderOptions
	Shadow   *MediaPosterShadowOptions
}

func (o MediaPosterBackgroundOptions) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *MediaPosterBackgroundOptions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to type assert MediaPosterBackgroundOptions")
	}
	return json.Unmarshal(bytes, &o)
}

type PosterItemPosition int

const (
	IconPositionMiddle PosterItemPosition = iota
	IconPositionTop
	IconPositionTopRight
	IconPositionTopLeft
	IconPositionLeft
	IconPositionRight
	IconPositionBottom
	IconPositionBottomLeft
	IconPositionBottomRight
)

type MediaPosterIconOptions struct {
	Enabled  *bool
	Path     *string
	Position *Coord `gorm:"type:json"`
	Size     *Size  `gorm:"type:json"`
}

func (o MediaPosterIconOptions) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *MediaPosterIconOptions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to type assert MediaPosterIconOptions")
	}
	return json.Unmarshal(bytes, &o)
}

type MediaPosterOverlayOptions struct {
	Enabled         bool
	Text            *string
	Transparency    *int
	CornerRadius    *int
	Position        *Coord                    `gorm:"type:json"`
	TextColor       *RGBAColor                `gorm:"type:json"`
	BackgroundColor *RGBAColor                `gorm:"type:json"`
	Icon            *MediaPosterIconOptions   `gorm:"type:json"`
	Border          *MediaPosterBorderOptions `gorm:"type:json"`
	Shadow          *MediaPosterShadowOptions `gorm:"type:json"`
}

func (o MediaPosterOverlayOptions) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *MediaPosterOverlayOptions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to type assert MediaPosterOverlayOptions")
	}
	return json.Unmarshal(bytes, &o)
}

type MediaImageType int

const (
	MediaImageTypeUnknown MediaImageType = iota
	MediaImageTypePoster
	MediaImageTypeBackground
	MediaImageTypeBanner
	MediaImageTypeLogo
	MediaImageTypeThumb
	MediaImageTypeClearArt
	MediaImageTypeDiscArt
)

type ProviderPoster struct {
	ID         string `gorm:"primaryKey"`
	ProviderId string
	Url        string
	Size       Size `gorm:"type:json"`
	Type       MediaImageType
}
