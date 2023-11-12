package main

import "testing"

//CalculateFlightRoute的单元测试
func TestCalculateFlightRoute(t *testing.T){
	flights := [][2]string{
		{"IND", "EWR"},
		{"SFO", "ATL"},
		{"GSO", "IND"},
		{"ATL", "GSO"},
	}
	cities := CalculateFlightRoute(flights)
	if cities[0] != "SFO" || cities[1] != "EWR" {
		t.Errorf("Expected be [SFO, EWR], but [%s, %s] got", cities[0], cities[1])
	}
}