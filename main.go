package main

import "fmt"

func CalculateFlightRoute(flights [][2]string) ([]string) {
	startSet := make(map[string]bool)
	endSet := make(map[string]bool)

	// 构建出发城市集合和到达城市集合
	for _, flight := range flights {
		startSet[flight[0]] = true
		endSet[flight[1]] = true
	}

	var startCity, endCity string
	var cities []string
	// 找到起始城市和终点城市，返回类型为数组，第一个值为起始城市，第二个值为终点城市
	for _, flight := range flights {
		if !endSet[flight[0]] {
			startCity = flight[0]
		}
		if !startSet[flight[1]] {
			endCity = flight[1]
		}
	}
	cities = append(cities, startCity, endCity)
	return cities
}

func main() {
	flights := [][2]string{
		{"IND", "EWR"},
    {"SFO", "ATL"},
    {"GSO", "IND"},
    {"ATL", "GSO"},
	}

	cities := CalculateFlightRoute(flights)
	fmt.Println(cities)
}