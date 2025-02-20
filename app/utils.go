package main

import (
	"fmt"
	"math"
	"os"

	"github.com/olekukonko/tablewriter"
)

func calculateDistance(lat1, lng1, lat2, lng2 float64) float64 {
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
	dist = dist * 60 * 1.1515 * 1.609344 // Convertir a kilómetros
	return dist
}

func printTable(properties []Property) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"price", "rooms", "bathrooms", "lighting", "location", "description"})

	for _, prop := range properties {
		location := fmt.Sprintf("%.4f, %.4f", prop.Location.Latitude, prop.Location.Longitude)
		desc := prop.Description
		if len(desc) > 27 {
			desc = desc[:27] + "..."
		}

		table.Append([]string{
			fmt.Sprintf("%.2f", prop.Price),
			fmt.Sprintf("%d", prop.Rooms),
			fmt.Sprintf("%d", prop.Bathrooms),
			string(prop.Lighting),
			location,
			desc,
		})
	}

	table.Render() // Renderiza la tabla en la terminal
}
