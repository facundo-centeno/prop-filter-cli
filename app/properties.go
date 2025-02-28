package main

import (
	"encoding/json"
	"os"
)

type LightingLevel string

const (
	Low    LightingLevel = "low"
	Medium LightingLevel = "medium"
	High   LightingLevel = "high"
)

func (ll LightingLevel) IsValid() bool {
	return ll == Low || ll == Medium || ll == High
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Property struct {
	SquareFootage int             `json:"square_footage"`
	Lighting      LightingLevel   `json:"lighting"`
	Price         float64         `json:"price"`
	Rooms         int             `json:"rooms"`
	Bathrooms     int             `json:"bathrooms"`
	Location      Location        `json:"location"`
	Description   string          `json:"description"`
	Amenities     map[string]bool `json:"amenities"`
}

func getProperties(fileName string) ([]Property, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var properties []Property
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&properties)
	if err != nil {
		return nil, err
	}

	return properties, nil
}
