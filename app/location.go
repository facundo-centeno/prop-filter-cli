package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GeoResponse struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func getUserLocation() (float64, float64, error) {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var geo GeoResponse
	if err := json.Unmarshal(body, &geo); err != nil {
		return 0, 0, err
	}

	fmt.Printf("Ubicación detectada: Latitud %.4f, Longitud %.4f\n", geo.Lat, geo.Lon)
	return geo.Lat, geo.Lon, nil
}
