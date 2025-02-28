package main

import (
	"flag"
	"fmt"
)

func main() {
	maxDistance := flag.Float64("max_distance", 0.0, "max distance from location")
	minPrice := flag.Float64("min_price", 1000, "Min price")
	maxPrice := flag.Float64("max_price", 1000000000, "Max price")
	minRooms := flag.Int("min_rooms", 1, "Min rooms")
	maxRooms := flag.Int("max_rooms", 200, "Max rooms")
	minBathrooms := flag.Int("min_bathrooms", 1, "Min bathrooms")
	maxBathrooms := flag.Int("max_bathrooms", 100, "Max bathrooms")
	lighting := flag.String("lighting_intensity", "", "Lighting intensity")
	description := flag.String("description", "", "Property description")
	amenities := flag.String("amenities", "", "Amenities")

	flag.Parse()

	userLat, userLon, err := getUserLocation()
	if err != nil {
		fmt.Println("Error al obtener la ubicación:", err)
		return
	}

	props, err := getProperties("../store/properties_collection.json")
	if err != nil {
		fmt.Println("Error al cargar propiedades:", err)
		return
	}

	filters := map[string]interface{}{
		"max_distance":       *maxDistance,
		"min_price":          *minPrice,
		"max_price":          *maxPrice,
		"min_rooms":          *minRooms,
		"max_rooms":          *maxRooms,
		"min_bathrooms":      *minBathrooms,
		"max_bathrooms":      *maxBathrooms,
		"lighting_intensity": LightingLevel(*lighting),
		"description":        *description,
		"amenities":          *amenities,
	}

	filteredProperties := filterProperties(props, filters, userLat, userLon)

	if len(filteredProperties) == 0 {
		fmt.Println("No hay propiedades dentro del rango especificado.")
		return
	}

	printTable(filteredProperties)
}
