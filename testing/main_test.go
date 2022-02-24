package main

import "testing"

func TestSum(t *testing.T) {
	// total := Sum(5, 5)

	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	// }

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{24, 12, 36},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{2, 4, 4},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d expected %d", max, item.n)
		}
	}
}

//go test -coverprofile=coverage.out
// ver resumen resumen en la terminal
//go tool cover -func=coverage.out
// o ver resumen en el navegador
//go tool cover -html=coverage.out