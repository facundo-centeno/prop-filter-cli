package main

import (
	"strings"
)

func filterProperties(properties []Property, filters map[string]any, userLat, userLon float64) []Property {
	var filtered []Property

	maxDistance, _ := filters["max_distance"].(float64)
	minPrice, _ := filters["min_price"].(float64)
	maxPrice, _ := filters["max_price"].(float64)
	minRooms, _ := filters["min_rooms"].(int)
	maxRooms, _ := filters["max_rooms"].(int)
	minBathrooms, _ := filters["min_bathrooms"].(int)
	maxBathrooms, _ := filters["max_bathrooms"].(int)
	lighting, _ := filters["lighting_intensity"].(LightingLevel)
	description, _ := filters["description"].(string)
	ammenities, _ := filters["ammenities"].(string)

	for _, property := range properties {
		if maxDistance > 0 {
			distance := calculateDistance(userLat, userLon, property.Location.Latitude, property.Location.Longitude)
			if distance > maxDistance {
				continue
			}
		}

		if (minPrice > 0 && property.Price < minPrice) ||
			(maxPrice > 0 && property.Price > maxPrice) {
			continue
		}

		if (minRooms > 0 && property.Rooms < minRooms) ||
			(maxRooms > 0 && property.Rooms > maxRooms) {
			continue
		}

		if (minBathrooms > 0 && property.Bathrooms < minBathrooms) ||
			(maxBathrooms > 0 && property.Bathrooms > maxBathrooms) {
			continue
		}

		if lighting.IsValid() && property.Lighting != lighting {
			continue
		}

		if description != "" &&
			!strings.Contains(strings.ToLower(property.Description), strings.ToLower(description)) {
			continue
		}

		if ammenities != "" && !property.Ammenities[ammenities] {
			continue
		}

		filtered = append(filtered, property)
	}

	return filtered
}
