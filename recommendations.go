package main

import (
	"fmt"
	"math"
)

var data = map[string]map[string]int{}

type Person struct {
	movies map[string]int
}

func main() {
	data["daniel"] = map[string]int{"test": 1, "four": 4}
	data["test"] = map[string]int{"test": 6, "four": 2, "ten": 10, "seven": 1}
	data["best"] = map[string]int{"test": 4, "four": 7, "ten": 4, "seven": 9}
	data["sest"] = map[string]int{"test": 2, "four": 1, "ten": 3, "seven": 3}

	fmt.Println(simDistance(data, "daniel", "test"))
	getRecommendations(data, "daniel")
}

func simDistance(data map[string]map[string]int, person1, person2 string) float64 {
	similiatiry := make(map[string]int)

	for item := range data[person1] {
		if _, ok := data[person2][item]; ok {
			similiatiry[item] = 1
		}
	}

	if len(similiatiry) == 0 {
		return 0.0
	}

	var sumOfSquares = 0.0
	for item := range data[person1] {
		if _, ok := data[person2][item]; ok {
			sumOfSquares += math.Pow(float64(data[person1][item])-float64(data[person2][item]), 2)
		}
	}

	return 1 / (1 + sumOfSquares)
}

func getRecommendations(data map[string]map[string]int, person string) float64 {
	totals := make(map[string]float64)
	simSums := make(map[string]float64)
	rankings := make(map[string]float64)
	for other := range data {
		if other == person {
			continue
		}

		sim := simDistance(data, person, other)
		if sim <= 0 {
			continue
		}
		for item := range data[other] {
			if _, ok := data[person]; !ok || data[person][item] == 0 {
				totals[item] += float64(data[other][item]) * sim
				simSums[item] += sim
			}
		}
	}

	for item, total := range totals {
		fmt.Println("total", total)
		fmt.Println("simSums", simSums[item])
		fmt.Println(float64(total / simSums[item]))
		rankings[item] = float64(total) / float64(simSums[item])
	}

	fmt.Println(rankings)

	return 0.0
}
