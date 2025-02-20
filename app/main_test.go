package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPriceFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low},
	}

	filters := map[string]interface{}{
		"min_price": 180000.0,
	}

	filtered := filterProperties(properties, filters, 0, 0)
	assert.Equal(t, 2, len(filtered))
}

func TestMaxPriceFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low},
	}

	filters := map[string]interface{}{
		"max_price": 400000.0,
	}

	filtered := filterProperties(properties, filters, 0, 0)
	assert.Equal(t, 2, len(filtered))
}

func TestMaxDistanceFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium,
			Location: Location{Latitude: 40.748817, Longitude: -73.985428}},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High,
			Location: Location{Latitude: 34.052235, Longitude: -118.243683}},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low,
			Location: Location{Latitude: 41.878113, Longitude: -87.629799}},
	}

	filters := map[string]interface{}{
		"max_distance": 400.0,
	}

	filtered := filterProperties(properties, filters, 40.748817, -73.985428)
	assert.Equal(t, 1, len(filtered))
}

func TestGetProperties(t *testing.T) {
	fileName := "test_properties.json"
	data := `[{"square_footage":1200,"lighting":"medium",
	"price":250000,"rooms":3,"bathrooms":2,
	"location":{"latitude":-31.427,"longitude":-64.189},
	"description":"cordoba","ammenities":{"garage":true,"pool":false}}]`

	err := os.WriteFile(fileName, []byte(data), 0644)
	assert.NoError(t, err)
	defer os.Remove(fileName)

	props, err := getProperties(fileName)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(props))
}

func TestLightingFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low},
	}

	filters := map[string]interface{}{
		"lighting_intensity": Medium,
	}

	filtered := filterProperties(properties, filters, 0, 0)
	assert.Equal(t, 1, len(filtered))
}

func TestDescriptionFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2,
			Lighting: Medium, Description: "great view"},
		{Price: 500000, Rooms: 5, Bathrooms: 4,
			Lighting: High, Description: "home with private pool"},
	}

	filters := map[string]interface{}{
		"description": "view",
	}

	filtered := filterProperties(properties, filters, 0, 0)
	assert.Equal(t, 1, len(filtered))
}

func TestAmmenitiesFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium,
			Ammenities: map[string]bool{"ping-pong": true}},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High,
			Ammenities: map[string]bool{"pool": true}},
	}

	filters := map[string]interface{}{
		"ammenities": "ping-pong",
	}

	filtered := filterProperties(properties, filters, 0, 0)
	assert.Equal(t, 1, len(filtered))
}
