package main

import (
	"fmt"
	"math"
	"sort"
)

var data = map[string]map[string]int{}

type Person struct {
	ID          int
	name        string
	movierating map[string]float64
}

func NewPerson(id int, name string, movierating map[string]float64) *Person {
	return &Person{
		ID:          id,
		name:        name,
		movierating: movierating,
	}
}

func (p *Person) AddMovieRating(name string, rating float64) {
	p.movierating[name] = rating
}

func (p *Person) GetMovies() map[string]float64 {
	return p.movierating
}

type PersonMap map[int]*Person

func NewPersonMap() PersonMap {
	m := make(PersonMap)
	return m
}

func (m PersonMap) GetPerson(id int) *Person {
	_, exists := m[id]
	if !exists {
		return nil
	}
	return m[id]
}

func (m PersonMap) GetPersonByPerson(person *Person) *Person {
	_, exists := m[person.ID]
	if !exists {
		return nil
	}
	return m[person.ID]
}

func (m PersonMap) AddPerson(person *Person) {
	m[person.ID] = NewPerson(person.ID, person.name, person.movierating)
}

func main() {

	person := Person{1, "daniel", map[string]float64{"test": 10, "four": 4, "star wars": 1, "five": 4}}
	//person := NewPerson(1, "daniel", map[string]float64{"test": 10, "four": 4, "star wars": 1, "five": 4})
	kalle := Person{2, "kalle", map[string]float64{"test": 6, "four": 2, "five": 6}}
	pelle := Person{3, "pelle", map[string]float64{"test": 1, "four": 4, "star wars": 10, "toy story": 4.5}}

	kalle.AddMovieRating("star wars", 6.5)
	kalle.AddMovieRating("six", 3.4)

	personMap := NewPersonMap()

	personMap.AddPerson(&person)
	personMap.AddPerson(&kalle)
	personMap.AddPerson(&pelle)

	fmt.Println(simDistance(personMap, &person, &kalle))
	fmt.Println(SimPearson(personMap, &person, &kalle))

	fmt.Println(getRecommendations(personMap, &kalle))

	//values := topMatches(personMap, person, 2)
}

func simDistance(data map[int]*Person, person1, person2 *Person) float64 {
	similiatiry := make(map[string]int)

	for movie := range data[person1.ID].GetMovies() {
		if _, ok := data[person2.ID].GetMovies()[movie]; ok {
			similiatiry[movie] = 1
		}
	}

	if len(similiatiry) == 0 {
		return 0.0
	}

	var sumOfSquares = 0.0
	for movie := range data[person1.ID].GetMovies() {
		if _, ok := data[person2.ID].GetMovies()[movie]; ok {
			sumOfSquares += math.Pow(float64(data[person1.ID].GetMovies()[movie])-float64(data[person2.ID].GetMovies()[movie]), 2)
		}
	}

	return 1 / (1 + sumOfSquares)
}

func SimPearson(data map[int]*Person, person1, person2 *Person) float64 {
	similiatiry := make(map[string]int)

	for movie := range data[person1.ID].GetMovies() {
		if _, ok := data[person2.ID].GetMovies()[movie]; ok {
			similiatiry[movie] = 1
		}
	}
	n := len(similiatiry)
	if n == 0 {
		return 0.0
	}

	sum1 := 0.0
	for movie := range similiatiry {
		sum1 += data[person1.ID].GetMovies()[movie]
	}
	sum2 := 0.0
	for movie := range similiatiry {
		sum2 += data[person2.ID].GetMovies()[movie]
	}

	sum1Sq := 0.0
	for movie := range similiatiry {
		sum1Sq += math.Pow(data[person1.ID].GetMovies()[movie], 2)
	}
	sum2Sq := 0.0
	for movie := range similiatiry {
		sum2Sq += math.Pow(data[person2.ID].GetMovies()[movie], 2)
	}

	pSum := 0.0
	for movie := range similiatiry {
		pSum += data[person2.ID].GetMovies()[movie] * data[person1.ID].GetMovies()[movie]
	}

	num := pSum - (sum1 * sum2 / 2)
	den := math.Sqrt((sum1Sq - math.Pow(sum1, 2)/float64(n)) * (sum2Sq - math.Pow(sum2, 2)/float64(n)))
	if den == 0 {
		return 0.0
	}

	r := num / den
	return r

	var sumOfSquares = 0.0
	for movie := range data[person1.ID].GetMovies() {
		if _, ok := data[person2.ID].GetMovies()[movie]; ok {
			sumOfSquares += math.Pow(float64(data[person1.ID].GetMovies()[movie])-float64(data[person2.ID].GetMovies()[movie]), 2)
		}
	}

	return 1 / (1 + sumOfSquares)
}

func topMatches(data map[int]*Person, person *Person, n int) []float64 {
	var scores []float64
	for _, other := range data {
		if other.ID == person.ID {
			continue
		}
		scores = append(scores, SimPearson(data, person, other))
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})
	return scores[:n]
}

func getRecommendations(data map[int]*Person, person *Person) PairList {
	totals := make(map[string]float64)
	simSums := make(map[string]float64)
	rankings := make(map[string]float64)
	for _, other := range data {
		if other.ID == person.ID {
			continue
		}
		sim := simDistance(data, person, other)
		if sim <= 0 {
			continue
		}

		for movie := range data[other.ID].GetMovies() {
			if _, ok := data[person.ID].GetMovies()[movie]; ok || data[person.ID].GetMovies()[movie] == 0 {
				totals[movie] += float64(data[other.ID].GetMovies()[movie]) * sim
				simSums[movie] += sim
			}
		}
	}

	for item, total := range totals {
		rankings[item] = float64(total) / float64(simSums[item])
	}

	sortedRankings := rankByValue(rankings)
	return sortedRankings
}

func rankByValue(rankings map[string]float64) PairList {
	pl := make(PairList, len(rankings))
	i := 0
	for k, v := range rankings {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TransformPreferences(data map[int]*Person) map[*Person]string {
	result := make(map[*Person]string)
	for _, person := range data {
		for movie := range data[person.ID].GetMovies() {
			if _, ok := data[person.ID].GetMovies()[movie]; ok {
				result[person] = movie
			}
		}
	}
	return result
}
