package main

type Square struct {
	X int
	Y int // confusing what with X and Os
}

type Triplet [3]Square

var triplets [8]Triplet


func fillTriplets() {
	triplets = [8]Triplet{
		Triplet{Square{0, 0}, Square{0, 1}, Square{0, 2}},
		Triplet{Square{1, 0}, Square{1, 1}, Square{1, 2}},
		Triplet{Square{2, 0}, Square{2, 1}, Square{2, 2}},
		Triplet{Square{0, 0}, Square{1, 0}, Square{2, 0}},
		Triplet{Square{0, 1}, Square{1, 1}, Square{2, 1}},
		Triplet{Square{0, 2}, Square{1, 2}, Square{2, 2}},
		Triplet{Square{0, 0}, Square{1, 1}, Square{2, 2}},
		Triplet{Square{0, 2}, Square{1, 1}, Square{2, 0}},
	}
}

func tripletsContaining(s Square) []Triplet{
	var searchResults []Triplet
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