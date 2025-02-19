package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
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

func (ll *LightingLevel) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed := LightingLevel(s)
	if !parsed.IsValid() {
		return fmt.Errorf("invalid lighting level: %s", s)
	}
	*ll = parsed
	return nil
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
	Ammenities    map[string]bool `json:"ammenities"`
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

func filterProps(properties []Property, filters map[string]any) []Property {
	var filtered []Property

	latitude, _ := filters["latitude"].(float64)
	longitude, _ := filters["longitude"].(float64)
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

		if minPrice > 0 && property.Price < minPrice ||
			maxPrice > 0 && property.Price > maxPrice {
			continue
		}

		if minRooms > 0 && property.Rooms < minRooms ||
			maxRooms > 0 && property.Rooms > maxRooms {
			continue
		}

		if minBathrooms > 0 && property.Bathrooms < minBathrooms ||
			maxBathrooms > 0 && property.Bathrooms > maxBathrooms {
			continue
		}

		if lighting.IsValid() && property.Lighting != lighting {
			continue
		}
		if maxDistance > 0 {
			dist := distance(latitude, longitude, property.Location.Latitude, property.Location.Longitude)
			if dist > maxDistance {
				continue
			}
		}
		if description != "" &&
			!strings.Contains(strings.ToLower(property.Description),
				strings.ToLower(description)) {
			continue
		}
		if ammenities != "" && !property.Ammenities[ammenities] {
			continue
		}

		filtered = append(filtered, property)
	}

	return filtered
}

func distance(lat1, lng1, lat2, lng2 float64) float64 {
	radlat1 := math.Pi * lat1 / 180
	radlat2 := math.Pi * lat2 / 180
	theta := lng1 - lng2
	radtheta := math.Pi * theta / 180
	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}
	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515 * 1.609344 // Convert to kilometers
	return dist
}

func main() {
	// for building executable  -->  go build -o prop-filter-cli
	// if permission denied --> chmod +x prop-filter-cli

	// query -->  ./prop-filter-cli -min_price 100000 -max_price 500000 -min_rooms 2 -max_rooms 5

	minSquareFootage := flag.Int("min_square_footage", 0, "min-square-footage")
	maxSquareFootage := flag.Int("max_square_footage", 100000, "max-square-footage")
	lighting := flag.String("lighting_intensity", "", "lighting-intensity")
	minPrice := flag.Float64("min_price", 1000, "min-price")
	maxPrice := flag.Float64("max_price", 1000000000, "max-price")

	maxRooms := flag.Int("max_rooms", 200, "max-rooms")
	minRooms := flag.Int("min_rooms", 1, "min-rooms")
	maxBathrooms := flag.Int("max_bathrooms", 100, "max-bathrooms")
	minBathrooms := flag.Int("min_bathrooms", 1, "min-bathrooms")

	description := flag.String("description", "", "description")
	ammenities := flag.String("ammenities", "", "ammenities")

	flag.Parse()

	props, err := getProperties("../store/properties_collection.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	filters := map[string]interface{}{
		"min_square_footage": *minSquareFootage,
		"max_square_footage": *maxSquareFootage,
		"min_price":          *minPrice,
		"max_price":          *maxPrice,
		"min_rooms":          *minRooms,
		"max_rooms":          *maxRooms,
		"min_bathrooms":      *minBathrooms,
		"max_bathrooms":      *maxBathrooms,
		"lighting_intensity": LightingLevel(*lighting),
		"description":        *description,
		"ammenities":         *ammenities,
	}

	filteredProperties := filterProps(props, filters)

	if len(filteredProperties) == 0 {
		fmt.Println("No properties match the given filters.")
		return
	}

	printTable(filteredProperties)

}

func printTable(properties []Property) {
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Printf("%-12s %-6s %-10s %-10s %-20s %-30s\n", "price", "rooms", "bathrooms", "lighting", "location", "description")
	fmt.Println("-----------------------------------------------------------------------------------")

	for _, prop := range properties {
		location := fmt.Sprintf("%.4f, %.4f", prop.Location.Latitude, prop.Location.Longitude)
		desc := prop.Description
		if len(desc) > 27 {
			desc = desc[:27] + "..."
		}
		fmt.Printf("%-12.2f %-6d %-10d %-10s %-20s %-30s\n", prop.Price, prop.Rooms, prop.Bathrooms, string(prop.Lighting), location, desc)
	}
	fmt.Println("-----------------------------------------------------------------------------------")
}
