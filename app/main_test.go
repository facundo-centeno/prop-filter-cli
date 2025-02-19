package main

import (
	"testing"
)

func TestMinPriceFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low},
	}

	filters := map[string]any{
		"min_price": 180000.0,
	}

	filtered := filterProps(properties, filters)

	if len(filtered) != 2 {
		t.Fatalf("expected 2 properties, got %d", len(filtered))
	}
}

func TestMaxPriceFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low},
	}

	filters := map[string]any{
		"max_price": 400000.0,
	}
	filtered := filterProps(properties, filters)

	if len(filtered) != 2 {
		t.Fatalf("expected 1 property, got %d", len(filtered))
	}
}

func TestMinRoomsFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low},
	}

	minRooms := 3
	filters := map[string]any{
		"min_rooms": minRooms,
	}
	filtered := filterProps(properties, filters)

	if len(filtered) != 2 {
		t.Fatalf("expected 2 properties, got %d", len(filtered))
	}
}

func TestMaxRoomsFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low},
	}

	maxRooms := 3
	filters := map[string]any{
		"max_rooms": maxRooms,
	}
	filtered := filterProps(properties, filters)

	if len(filtered) != 2 {
		t.Fatalf("expected 2 properties, got %d", len(filtered))
	}
}

func TestMinBathroomsFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low},
	}

	minBathrooms := 2
	filters := map[string]any{
		"min_bathrooms": minBathrooms,
	}
	filtered := filterProps(properties, filters)

	if len(filtered) != 2 {
		t.Fatalf("expected 2 properties, got %d", len(filtered))
	}
}

func TestMaxBathroomsFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low},
	}

	maxBathrooms := 2
	filters := map[string]any{
		"max_bathrooms": maxBathrooms,
	}
	filtered := filterProps(properties, filters)

	if len(filtered) != 2 {
		t.Fatalf("expected 2 properties, got %d", len(filtered))
	}
}

func TestLightingIntensityFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low},
	}

	lighting := "high"
	filters := map[string]any{
		"lighting_intensity": LightingLevel(lighting),
	}
	filtered := filterProps(properties, filters)

	if len(filtered) != 1 {
		t.Fatalf("expected 1 property, got %d", len(filtered))
	}
}

func TestDescriptionFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium, Description: "big apartment"},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High, Description: "small aptmnt"},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low, Description: "beautiful view"},
	}

	description := "view"
	filters := map[string]any{
		"description": description,
	}
	filtered := filterProps(properties, filters)

	if len(filtered) != 1 {
		t.Fatalf("expected 1 property, got %d", len(filtered))
	}
}

func TestAmmenitiesFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium, Ammenities: map[string]bool{"pool": true}},
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High, Ammenities: map[string]bool{"gym": true}},
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low, Ammenities: map[string]bool{"ping-pong-table": true}},
	}

	filters := map[string]any{
		"ammenities": "pool",
	}
	filtered := filterProps(properties, filters)

	if len(filtered) != 1 {
		t.Fatalf("expected 1 property, got %d", len(filtered))
	}
}

func TestMaxDistanceFilter(t *testing.T) {
	properties := []Property{
		{Price: 200000, Rooms: 3, Bathrooms: 2, Lighting: Medium, Location: Location{Latitude: 40.748817, Longitude: -73.985428}}, // coordinates of NY City
		{Price: 500000, Rooms: 5, Bathrooms: 4, Lighting: High, Location: Location{Latitude: 34.052235, Longitude: -118.243683}},  // coordinates of Los Angeles
		{Price: 150000, Rooms: 2, Bathrooms: 1, Lighting: Low, Location: Location{Latitude: 41.878113, Longitude: -87.629799}},    // coordinates of Chicago
	}

	filters := map[string]any{
		"latitude":     40.748817,  // NY latitude
		"longitude":    -73.985428, // NY Longitude
		"max_distance": 400.0,      // max distance
	}

	filtered := filterProps(properties, filters)

	// we expect only one prop inside the 400km ratio
	if len(filtered) != 1 {
		t.Fatalf("expected 1 property, got %d", len(filtered))
	}
}
