package types

import "github.com/google/uuid"

type ClientType int

const (
	ClientTypeUnknown ClientType = iota
	ClientTypeMediaServer
	ClientTypeListProvider
	ClientTypeUtility
)

type FieldType int

const (
	FieldTypeString FieldType = iota
	FieldTypeBool
	FieldTypeNumber
	FieldTypePassword
)

type ClientField struct {
	ID          string
	ClientId    string
	Name        string
	Placeholder string
	Type        FieldType
}

type Client struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name           string
	Label          string
	ConfigClients  []*ConfiguredClient `gorm:"foreignKey:ClientID"`
	LibraryClients []*LibraryClient    `gorm:"foreignKey:ClientID"`
	Type           ClientType
}

type ConfiguredClientFieldValues struct {
	ID                      uuid.UUID `gorm:"type:uuid;primaryKey;"`
	ConfiguredClientFieldID uuid.UUID
	ConfiguredClientID      uuid.UUID
	Value                   string
	ClientField             ClientField `gorm:"foreignKey:ConfiguredClientFieldID"`
}

type ConfiguredClient struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Label             string
	Client            *Client `gorm:"foreignKey:ID"`
	ClientID          uuid.UUID
	ConfigID          uuid.UUID
	ClientFields      []*ClientField                 `gorm:"foreignKey:ID"`
	ClientFieldValues []*ConfiguredClientFieldValues `gorm:"foreignKey:ID"`
}
