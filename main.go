package main

import (
    "encoding/json"

    "math"
    "net/http"
    "strconv"
)

const earthRadiusKm = 6371.0

type DistanceResponse struct {
    Distance float64 `json:"distance"` // Mesafe kilometre cinsinden
}

func main() {
    http.HandleFunc("/distance", distanceHandler)
    http.ListenAndServe(":8080", nil)
}

func distanceHandler(w http.ResponseWriter, r *http.Request) {
    lat1Str := r.URL.Query().Get("lat1")
    lon1Str := r.URL.Query().Get("lon1")
    lat2Str := r.URL.Query().Get("lat2")
    lon2Str := r.URL.Query().Get("lon2")

    if lat1Str == "" || lon1Str == "" || lat2Str == "" || lon2Str == "" {
        http.Error(w, "Missing coordinates", http.StatusBadRequest)
        return
    }

    lat1, err := strconv.ParseFloat(lat1Str, 64)
    if err != nil {
        http.Error(w, "Invalid latitude for point A", http.StatusBadRequest)
        return
    }
    lon1, err := strconv.ParseFloat(lon1Str, 64)
    if err != nil {
        http.Error(w, "Invalid longitude for point A", http.StatusBadRequest)
        return
    }
    lat2, err := strconv.ParseFloat(lat2Str, 64)
    if err != nil {
        http.Error(w, "Invalid latitude for point B", http.StatusBadRequest)
        return
    }
    lon2, err := strconv.ParseFloat(lon2Str, 64)
    if err != nil {
        http.Error(w, "Invalid longitude for point B", http.StatusBadRequest)
        return
    }

    distance := haversine(lat1, lon1, lat2, lon2)

    response := DistanceResponse{
        Distance: distance,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
    dLat := degreesToRadians(lat2 - lat1)
    dLon := degreesToRadians(lon2 - lon1)

    a := math.Sin(dLat/2)*math.Sin(dLat/2) +
        math.Cos(degreesToRadians(lat1))*math.Cos(degreesToRadians(lat2))*
            math.Sin(dLon/2)*math.Sin(dLon/2)

    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

    return earthRadiusKm * c
}

func degreesToRadians(degrees float64) float64 {
    return degrees * math.Pi / 180
}
