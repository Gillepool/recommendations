package recommendations

import (
	"fmt"
	"testing"
)

func TestPerson(t *testing.T) {

	person := Person{1, "daniel", map[string]float64{"test": 10, "four": 4, "star wars": 1, "five": 4}}
	person2 := Person{2, "whodis", map[string]float64{"test": 6, "four": 2, "five": 6}}
	person3 := Person{3, "random guy", map[string]float64{"test": 1, "four": 4, "star wars": 10, "toy story": 4.5}}

	person2.AddMovieRating("star wars", 6.5)
	person2.AddMovieRating("six", 3.4)

	personMap := NewPersonMap()
	personMap.AddPerson(&person)
	personMap.AddPerson(&person2)
	personMap.AddPerson(&person3)

	if len(personMap) < 3 {
		t.Error("Error, not everyone was added to the map")
	}
}

func TestRecommendations(t *testing.T) {
	person := Person{1, "daniel", map[string]float64{"test": 10, "four": 4, "star wars": 1, "five": 4}}
	person2 := Person{2, "whodis", map[string]float64{"test": 6, "four": 2, "five": 6}}
	person3 := Person{3, "random guy", map[string]float64{"test": 1, "four": 4, "star wars": 10, "toy story": 4.5}}

	person2.AddMovieRating("star wars", 6.5)
	person2.AddMovieRating("six", 3.4)

	personMap := NewPersonMap()
	personMap.AddPerson(&person)
	personMap.AddPerson(&person2)
	personMap.AddPerson(&person3)
	recommendations := GetRecommendations(personMap, &person)
	fmt.Println(recommendations)
}
