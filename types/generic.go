package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type RGBAColor struct {
	R int
	G int
	B int
	A int
}

func (c RGBAColor) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *RGBAColor) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion for RGBAColor failed")
	}
	return json.Unmarshal(bytes, &c)
}

type Size struct {
	W int
	H int
}

func (s Size) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Size) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion for Size failed")
	}
	return json.Unmarshal(bytes, &s)
}

type Coord struct {
	X int
	Y int
}

func (c Coord) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Coord) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion for Coord failed")
	}
	return json.Unmarshal(bytes, &c)
}
