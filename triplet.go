package main

type Square struct {
	X int
	Y int // confusing what with X and Os
}

type TripletPosition [3]Square

var triplets [8]TripletPosition


func defineTripletPositions() {
	triplets = [8]TripletPosition{
		TripletPosition{Square{0, 0}, Square{0, 1}, Square{0, 2}},
		TripletPosition{Square{1, 0}, Square{1, 1}, Square{1, 2}},
		TripletPosition{Square{2, 0}, Square{2, 1}, Square{2, 2}},
		TripletPosition{Square{0, 0}, Square{1, 0}, Square{2, 0}},
		TripletPosition{Square{0, 1}, Square{1, 1}, Square{2, 1}},
		TripletPosition{Square{0, 2}, Square{1, 2}, Square{2, 2}},
		TripletPosition{Square{0, 0}, Square{1, 1}, Square{2, 2}},
		TripletPosition{Square{0, 2}, Square{1, 1}, Square{2, 0}},
	}
}

func tripletsContaining(s Square) []TripletPosition{
	var searchResults []TripletPosition
	for _, t := range triplets {
		for _, squareInTriplet := range t {
			if s == squareInTriplet {
				searchResults = append(searchResults, t)
				break
			}
		}
	}
	return searchResults
}